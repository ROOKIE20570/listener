package parser

import (
	"io"
	"strconv"
)

type Mysql struct {
}

func (mysql *Mysql) Resolve(r io.Reader) (string, error) {
	return "", nil
}

func (mysql *Mysql) GetFilter(port int) string {
	return "tcp and port " + strconv.Itoa(port)
}
