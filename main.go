package main

import (
	"github.com/cboy868/china_regions/engine"
	"github.com/cboy868/china_regions/region/models"
	"github.com/cboy868/china_regions/region/parser"
)

func main() {

	url := "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2019/index.html"
	// url := "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2019/53/01/02/530102001.html"
	engine.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseProList,
		Pitem:      models.Region{Code: "0"},
	})
}
