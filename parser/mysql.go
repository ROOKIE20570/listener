package parser

import "strconv"

type Mysql struct {
}

func (mysql *Mysql) Resolve(stream []byte) (string, error) {
	return "", nil
}

func (mysql *Mysql) GetFilter(port int) string {
	return "tcp and port " + strconv.Itoa(port)
}
