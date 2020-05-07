package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

//Fetch 抓取数据
func Fetch(url string) ([]byte, error) {
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil) //建立一个请求
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}
	//Add 头协议
	reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	reqest.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	reqest.Header.Add("Connection", "keep-alive")
	reqest.Header.Add("Cookie", "_trs_uv=k9jf5tla_6_bq96; AD_RS_COOKIE=20080917; wzws_cid=50297ada8d5ff3de9372f177d9a12cba30a4556b2b6dc06b8706d7d2f70647a0677de6e09cfe2a6ceab6b67a5f5970f495c686addac20b2c6348d4efb78cebc1cf35756302e81f0b75abc8ce6495bfaa")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36")
	resp, err := client.Do(reqest) //提交

	// resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error status Code:%d", resp.StatusCode)
	}

	reader := bufio.NewReader(resp.Body)
	e := determineEncoding(reader)
	utf8Reader := transform.NewReader(reader, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

//尝试转码
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("error encoding:%v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
