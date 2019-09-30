package parser

import "strconv"

type Mysql struct {
}

func (mysql *Mysql) resolve(stream []byte) (string, error) {
	return "", nil
}

func (mysql *Mysql) getFilter(port int) string {
	return "tcp and port " + strconv.Itoa(port)
}
