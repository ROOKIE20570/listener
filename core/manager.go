package core

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/tcpassembly"
	"github.com/google/gopacket/tcpassembly/tcpreader"
	"listener/parser"
	"log"
	"time"
)

type pluginStreamFactory struct{}

type pluginStream struct {
	net, transport gopacket.Flow
	r              tcpreader.ReaderStream
}

var parse parser.Parser

func Run(device, listener string, port int) {

	handle, err := pcap.OpenLive(device, 65535, false, pcap.BlockForever)
	if err != nil {
		log.Fatal("open listener fail", err)
	}
	switch listener {
	case "mysql":
		parse = new(parser.Mysql)
	case "redis":
		parse = new(parser.Redis)
	default:
		log.Fatal("not supported")
	}
	handle.SetBPFFilter(parse.GetFilter(port))

	streamFactory := &pluginStreamFactory{}
	streamPool := tcpassembly.NewStreamPool(streamFactory)
	assembler := tcpassembly.NewAssembler(streamPool)

	flushTicker := time.Tick(1 * time.Minute)
	log.Printf("listening device %s,port %d, type %s", device, port, listener)
	packets := gopacket.NewPacketSource(handle, handle.LinkType()).Packets()

	for {
		select {
		case packet := <-packets:
			if packet.NetworkLayer() == nil || packet.TransportLayer() == nil || packet.TransportLayer().LayerType() != layers.LayerTypeTCP {
				continue
			}
			tcpLayer := packet.TransportLayer().(*layers.TCP)
			assembler.AssembleWithTimestamp(packet.NetworkLayer().NetworkFlow(), tcpLayer, packet.Metadata().Timestamp)

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
	go // Important... we must guarantee that data from the reader stream is read.

	// ReaderStream implements tcpassembly.Stream, so we can return a pointer to it.
	return &pluginStream.r
}
