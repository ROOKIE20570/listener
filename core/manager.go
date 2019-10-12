package core

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"listener/parser"
	"log"
)

func Run(device, listener string, port int) {

	handle, err := pcap.OpenLive(device, 65535, false, pcap.BlockForever)
	if err != nil {
		log.Fatal("open listener fail", err)
	}
	var parse parser.Parser
	switch listener {
	case "mysql":
		parse = new(parser.Mysql)
	case "redis":
		parse = new(parser.Redis)
	default:
		log.Fatal("not supported")
	}
	log.Println(parse.GetFilter(port))
	handle.SetBPFFilter(parse.GetFilter(port))

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		tcpLayer := packet.Layer(layers.LayerTypeTCP)
		if tcpLayer != nil {
			tcp, _ := tcpLayer.(*layers.TCP)
		}
	}
}
