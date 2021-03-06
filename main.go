package main

import (
	"github.com/cboy868/china_regions/engine"
	"github.com/cboy868/china_regions/region/models"
	"github.com/cboy868/china_regions/region/parser"
	"github.com/cboy868/china_regions/scheduler"
)

func main() {
	// provinceURL := "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2019/index.html"
	// // url := "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2019/53/01/02/530102001.html"
	// engine.SimpleEngine{}.Run(engine.Request{
	// 	Url:        provinceURL,
	// 	ParserFunc: parser.ParseProList,
	// 	Pitem:      models.Region{Code: "0", Type: "provincetr"},
	// })

	// cityURL := "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2019/13.html"
	// engine.Run(engine.Request{
	// 	Url:        cityURL,
	// 	ParserFunc: parser.ParseCityList,
	// 	Pitem:      models.Region{Code: "0", Type: "citytr"},
	// })

	provinceURL := "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2019/index.html"
	// url := "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2019/53/01/02/530102001.html"
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url:        provinceURL,
		ParserFunc: parser.ParseProList,
		Pitem:      models.Region{Code: "0", Type: "provincetr"},
	})
}
