package core

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"listener/parser"
	"log"
)
func Run(device string, port int){

	handle, err := pcap.OpenLive(device, 65535, false, 0)
	if err != nil{
		log.Fatal("open listener fail")
	}
	var parse parser.Parser
	switch device {
	case "mysql":
		parse = new(parser.Mysql)
	case "redis":
		parse = new(parser.Redis)
	default:
		log.Fatal("not supported")
	}

	handle.SetBPFFilter(parse.GetFilter(port))

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet:=range packetSource.Packets(){
		fmt.Println(packet)
	}
}