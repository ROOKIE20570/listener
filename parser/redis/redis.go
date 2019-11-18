package redis

import (
	"io"
	"strconv"
)

type Redis struct {
}

func (redis *Redis) Resolve(r io.Reader) (string, error) {
	return "", nil
}

func (redis *Redis) GetFilter(port int) string {
	return "tcp and port " + strconv.Itoa(port)
}
