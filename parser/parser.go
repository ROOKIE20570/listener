package parser

import "io"

type Parser interface {
	//对应每个parser不同的过滤行为
	Resolve(r io.Reader) (string, error) //解析函数
	GetFilter(port int) string
}
