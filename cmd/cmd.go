package main

import (
	"flag"
	"listener/core"
)

const version = "0.0.0"

var (
	port     int
	device   string
	listener string
)

func main() {
	//todo 参数校验 提示
	flag.IntVar(&port, "p", 0, "port number")
	flag.StringVar(&device, "d", "", "device name")
	flag.StringVar(&listener, "l", "", "listener name")

	flag.Parse()
	core.Run(device, listener, port)

}


