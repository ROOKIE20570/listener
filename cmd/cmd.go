package main

import (
	"flag"
	"listener/core"
)

const version  = "0.0.0"
var (
	port int
	device string
)

func main() {
	flag.IntVar(&port,"p",0,"port number")
	flag.StringVar(&device,"d","","device name")

	core.Run(device,port)


}
