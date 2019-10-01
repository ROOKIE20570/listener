package parser

type parser interface {
	 //对应每个parser不同的过滤行为
	resolve (stream []byte) (string, error)//解析函数
	getFilter(port int) string
}
