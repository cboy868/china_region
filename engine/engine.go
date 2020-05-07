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

		// f, err := os.Create("./region.sql") //创建文件
		// if err != nil {
		// 	log.Printf("打开文件错误：%s", err)
		// }
		// defer f.Close()

		client := CreateRedisClient()
		defer client.Close()

		for _, item := range parseResult.Items {
			log.Printf("Got items: %v", item)
			// fmt.Printf("a:%s", f)
			// WriteSQL(item.(models.Region), f)
			WriteSQLToRedis(item.(models.Region), client)
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

//WriteSQLToRedis 2
func WriteSQLToRedis(region models.Region, client *redis.Client) {
	sqlStrFmt := "INSERT INTO region (name, pcode, code) VALUES ('%s','%s', '%s');\n"
	sqlStr := fmt.Sprintf(sqlStrFmt, region.Name, region.Pcode, region.Code)
	client.HMSet("myredis", region.Code, sqlStr)
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
