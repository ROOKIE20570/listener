package parser

type parser struct {
	filter  string //对应每个parser不同的过滤行为
	resolve func([]byte) (string, error)//解析函数

}
