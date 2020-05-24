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

//Worker 1
func Worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		return ParseResult{}, err
	}
	return r.ParserFunc(r.Url, body, r.Pitem), nil
}

// WriteCsvFile 数据写入csv文件
func WriteCsvFile(region models.Region, f *os.File) {
	w := bufio.NewWriterSize(f, 4096) //创建新的 Writer 对象
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

//OpenFiles 1
func OpenFiles() map[string]*os.File {
	c := make(map[string]*os.File)
	s := []string{"provincetr", "citytr", "countytr", "towntr", "villagetr"}
	for _, city := range s {
		f, err := os.OpenFile("./files/"+city+".csv", os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Printf("error:%s", err)
		}
		c[city] = f
	}
	return c
}
