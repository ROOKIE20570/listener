package core

import (
	"github.com/google/gopacket/pcap"
	"listener/parser"
	"log"
)
func run(device string, port int){

	handle, err := pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	var parse *parser.Parser
	switch device {
	case "mysql":
		parse = new(parser.Mysql)
	case "redis":
		parse = new(parser.Redis)
	default:
		log.Fatal("not supported")
	}

	handle.setBPFFilter(parse.getFilter())
}