package engine

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serializ() (name string, args interface{})
}

type Request struct {
	Url    string
	Parser Parser
}

type ParserFunc func(contents []byte, url string) ParseResult

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Id      string
	Type    string
	Url     string
	PayLoad interface{}
}

type NilParser struct {
}

func (NilParser) Serializ() (name string, args interface{}) {
	return "NilParser", nil
}

func (NilParser) Parse(b []byte, url string) ParseResult {
	return ParseResult{}
}

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serializ() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
