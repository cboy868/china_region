package main

import (
	"github.com/cboy868/china_regions/engine"
	"github.com/cboy868/china_regions/region/models"
	"github.com/cboy868/china_regions/region/parser"
)

func main() {

	// engine.WriteSQL(models.Region{
	// 	Code:  "1",
	// 	Name:  "name1",
	// 	Pcode: "0",
	// })

	// return

	// s := "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2019/index.html"
	// //解析这个 URL 并确保解析没有出错。

	// fmt.Printf("url:%s", s[0:strings.LastIndex(s, "/")])

	// url := "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2019/index.html"
	url := "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2019/53/01/02/530102001.html"

	// url := "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2019/13/1301.html"
	engine.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseVillageList,
		Pitem:      models.Region{Code: "0"},
	})
	// engine.Run(engine.Request{
	// 	Url:        url,
	// 	ParserFunc: parser.ParseProList,
	// 	Pitem:      models.Region{Code: "0"},
	// })
}
