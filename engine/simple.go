package engine

import (
	"log"

	"github.com/cboy868/china_regions/region/models"
)

//SimpleEngine 122
type SimpleEngine struct{}

//Run start
func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		// body, err := fetcher.Fetch(r.Url)

		// log.Printf("Get data from url:%s", r.Url)

		// if err != nil {
		// 	log.Printf("fetch url:%v error:%v 重新把请求推入到requests中", r.Url, err)
		// 	requests = append(requests, r)
		// 	continue
		// }

		// // log.Printf("内容：%s", body)
		// // cityStr := (r.Pitem.(models.Region)).Type

		// parseResult := r.ParserFunc(r.Url, body, r.Pitem)

		parseResult, err := Worker(r)
		if err != nil {
			log.Printf("fetch url:%v error:%v 重新把请求推入到requests中", r.Url, err)
			requests = append(requests, r)
			continue
		}

		requests = append(requests, parseResult.Requests...)

		// f, err := OpenFile("./files/" + cityStr + ".csv")
		// if err != nil {
		// 	panic("有错先直接异常" + err.Error())
		// }
		// defer f.Close()

		fileMaps := OpenFiles()

		// defer f.Close()

		for _, item := range parseResult.Items {
			log.Printf("Got items: %v", item)
			f := fileMaps[(item.(models.Region)).Type]
			defer f.Close()
			WriteCsvFile(item.(models.Region), f)
		}
	}

}
