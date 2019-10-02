package main

var (
	port int
	name string
)

func main() {

	////Open device
	//handle, err := pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer handle.Close()
	//
	//// Use the handle as a packet source to process all packets
	//packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	//for packet := range packetSource.Packets() {
	//	// Process packet here
	//	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	//	if tcpLayer != nil{
	//		_, some := tcpLayer.(*layers.TCP)
	//		fmt.Println(some)
	//	}
	//}


}
