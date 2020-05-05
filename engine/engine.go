package engine

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"xue.com/wansq/china_regions/fetcher"
	"xue.com/wansq/china_regions/region/models"
)

//Run start
func Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		body, err := fetcher.Fetch(r.Url)

		log.Printf("Get data from url:%s", r.Url)

		if err != nil {
			log.Printf("fetch url:%v error:%v 重新把请求推入到requests中", r.Url, err)
			requests = append(requests, r)
			continue
		}

		// log.Printf("内容：%s", body)

		parseResult := r.ParserFunc(r.Url, body, r.Pitem)

		requests = append(requests, parseResult.Requests...)

		f, err := os.Create("./region.sql") //创建文件
		if err != nil {
			log.Printf("打开文件错误：%s", err)
		}
		defer f.Close()

		for _, item := range parseResult.Items {
			log.Printf("Got items: %v", item)
			// fmt.Printf("a:%s", f)
			WriteSQL(item.(models.Region), f)
		}
	}

}

//WriteSQL 1
func WriteSQL(region models.Region, f *os.File) {
	sqlStrFmt := "INSERT INTO region (name, pcode, code) VALUES ('%s','%s', '%s');\n"
	sqlStr := fmt.Sprintf(sqlStrFmt, region.Name, region.Pcode, region.Code)

	w := bufio.NewWriter(f) //创建新的 Writer 对象
	_, err := w.WriteString(sqlStr)
	if err != nil {
		log.Printf("打开文件错误：%s", err)
	}
	w.Flush()
}
