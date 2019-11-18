package parser

import (
	"github.com/google/gopacket"
	"io"
)

type Parser interface {
	//对应每个parser不同的过滤行为
	Resolve(net gopacket.Flow, transport gopacket.Flow, r io.Reader) (string, error) //解析函数
	GetFilter(port int) string
}
