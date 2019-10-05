package parser

import "strconv"

type Redis struct {
}

func (redis *Redis) Resolve(stream []byte) (string, error) {

}

func (redis *Redis) GetFilter(port int) string {
	return "tcp and port " + strconv.Itoa(port)
}
