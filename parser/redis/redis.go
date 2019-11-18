package redis

import (
	"github.com/google/gopacket"
	"io"
	"strconv"
)

type Redis struct {
}

func (redis *Redis) Resolve(net gopacket.Flow, transport gopacket.Flow, r io.Reader) (string, error) {
	return "", nil
}

func (redis *Redis) GetFilter(port int) string {
	return "tcp and port " + strconv.Itoa(port)
}
