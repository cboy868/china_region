package engine

//Request 请求
type Request struct {
	Url        string
	Pitem      interface{}
	ParserFunc func(string, []byte, interface{}) ParseResult
}

//ParseResult 解析结果
type ParseResult struct {
	Items    []interface{}
	Requests []Request
}

//NilParser 空解析
func NilParser(string, []byte, interface{}) ParseResult {
	return ParseResult{}
}
