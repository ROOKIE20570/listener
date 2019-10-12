package parser

type parser interface {
	FilterString () string
	ResolveStream([]byte)(string,error)
}