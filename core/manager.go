package core

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/tcpassembly"
	"github.com/google/gopacket/tcpassembly/tcpreader"
	"listener/config"
	"listener/parser"
	"listener/parser/mysql"
	"listener/parser/redis"
	"log"
	"time"
)

type pluginStreamFactory struct{}

type pluginStream struct {
	net, transport gopacket.Flow
	r              tcpreader.ReaderStream
}

var parse parser.Parser

func Run() {

	handle, err := pcap.OpenLive(config.Conf.Device, 65535, false, pcap.BlockForever)
	if err != nil {
		log.Fatal("open listener fail", err)
	}
	switch config.Conf.Type {
	case "mysql":
		parse = new(mysql.Mysql)
	case "redis":
		parse = new(redis.Redis)
	default:
		log.Fatal("not supported")
	}

	err = handle.SetBPFFilter(parse.GetFilter(config.Conf.Port))
	if err != nil {
		log.Println("set filter fail")
		panic(err)
	}
	streamFactory := &pluginStreamFactory{}
	streamPool := tcpassembly.NewStreamPool(streamFactory)
	assembler := tcpassembly.NewAssembler(streamPool)

	flushTicker := time.Tick(1 * time.Minute)
	log.Printf("listening device %s,port %d, type %s", config.Conf.Device, config.Conf.Port, config.Conf.Type)
	packets := gopacket.NewPacketSource(handle, handle.LinkType()).Packets()

	for {
		select {
		case packet := <-packets:
			if packet == nil {
				return
			}

			if packet.NetworkLayer() == nil || packet.TransportLayer() == nil || packet.TransportLayer().LayerType() != layers.LayerTypeTCP {
				//Unusable packet
				continue
			}

			tcp := packet.TransportLayer().(*layers.TCP)
			assembler.AssembleWithTimestamp(packet.NetworkLayer().NetworkFlow(), tcp, packet.Metadata().Timestamp)

		case <-flushTicker:
			assembler.FlushOlderThan(time.Now().Add(time.Minute * -2))
		}
	}
}

func (pluginStreamFactory *pluginStreamFactory) New(net, transport gopacket.Flow) tcpassembly.Stream {
	pluginStream := &pluginStream{
		net:       net,
		transport: transport,
		r:         tcpreader.NewReaderStream(),
	}

	go parse.Resolve(net, transport, &(pluginStream.r))
	// Important... we must guarantee that data from the reader stream is read.
	// ReaderStream implements tcpassembly.Stream, so we can return a pointer to it.
	return &pluginStream.r
}
