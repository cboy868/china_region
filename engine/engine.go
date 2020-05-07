package engine

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/cboy868/china_regions/fetcher"
	"github.com/cboy868/china_regions/region/models"
	"github.com/go-redis/redis/v7"
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

		cityStr := (r.Pitem.(models.Region)).Type
		f, err := OpenFile("./files/" + cityStr + ".csv")
		if err != nil {
			panic("有错先直接异常" + err.Error())
		}
		defer f.Close()

		for _, item := range parseResult.Items {
			log.Printf("Got items: %v", item)
			WriteCsvFile(item.(models.Region), f)
		}
	}

}

// WriteCsvFile 数据写入csv文件
func WriteCsvFile(region models.Region, f *os.File) {
	w := bufio.NewWriter(f) //创建新的 Writer 对象
	lineStr := fmt.Sprintf("%s,%s,%s,%s\n", region.Code, region.Pcode, region.Name, region.Type)
	_, err := w.WriteString(lineStr)
	if err != nil {
		log.Printf("打开文件错误：%s", err)
	}
	w.Flush()
}

//WriteSQLToFile 1
func WriteSQLToFile(region models.Region, f *os.File) {
	sqlStrFmt := "INSERT INTO region (name, pcode, code) VALUES ('%s','%s', '%s');\n"
	sqlStr := fmt.Sprintf(sqlStrFmt, region.Name, region.Pcode, region.Code)

	w := bufio.NewWriter(f) //创建新的 Writer 对象
	_, err := w.WriteString(sqlStr)
	if err != nil {
		log.Printf("打开文件错误：%s", err)
	}
	w.Flush()
}

//WriteSQLToRedis 2
func WriteSQLToRedis(region models.Region, client *redis.Client) {
	sqlStrFmt := "INSERT INTO region (name, pcode, code) VALUES ('%s','%s', '%s');\n"
	sqlStr := fmt.Sprintf(sqlStrFmt, region.Name, region.Pcode, region.Code)
	client.HMSet("myredis", region.Code, sqlStr)
}

//OpenFile 获取文件句柄
func OpenFile(file string) (*os.File, error) {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("error:%s", err)
		return nil, err
	}
	return f, nil
}

//CreateRedisClient 1
func CreateRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}
