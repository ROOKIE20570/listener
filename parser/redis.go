package parser

import "strconv"

type Redis struct {
}

func (redis *Redis) resolve(stream []byte) (string, error) {

}

func (redis *Redis) getFilter(port int) string {
	return "tcp and port " + strconv.Itoa(port)
}
