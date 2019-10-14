package parser

type Parser interface {
	//对应每个parser不同的过滤行为
	Resolve(stream []byte) (string, error) //解析函数
	GetFilter(port int) string
}
