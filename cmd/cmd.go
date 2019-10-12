package main

import (
	"flag"
	"listener/core"
	"log"
)

const version = "0.0.0"

var (
	port     int
	device   string
	listener string
)

func main() {
	flag.IntVar(&port, "p", 0, "port number")
	flag.StringVar(&device, "d", "", "device name")
	flag.StringVar(&listener, "l", "", "listener name")

	flag.Parse()
	log.Println(device,listener,port)
	core.Run(device, listener, port)

}
