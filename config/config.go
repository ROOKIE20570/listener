package config

import "flag"

type Config struct {
	Type   string
	Device string
	Port   int
}

var Conf *Config

func init() {

	flag.IntVar(&Conf.Port, "p", 3306, "port number")
	flag.StringVar(&Conf.Device, "d", "lo0", "device name")
	flag.StringVar(&Conf.Type, "l", "mysql", "listen type")
	flag.Parse()
}

