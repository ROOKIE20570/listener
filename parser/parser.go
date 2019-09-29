package parser

type plugin struct {
	filter string
	port int
}

type parser interface {
	 //对应每个parser不同的过滤行为
	resolve ([]byte) (string, error)//解析函数
	newParser() *plugin


}
