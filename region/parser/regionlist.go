package parser

import (
	"regexp"
	"strings"

	"github.com/cboy868/china_regions/engine"
	"github.com/cboy868/china_regions/region/models"
)

const provinceListRe = `<td><a href='([0-9]+)\.html'>([^<]+)<br/></a></td>`

//ParseProList 1
func ParseProList(url string, contents []byte, pitem interface{}) engine.ParseResult {
	re := regexp.MustCompile(provinceListRe)

	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	// fmt.Println(url[0 : len(url)-10])
	// fmt.Printf("abc:%v", matches)
	// fmt.Printf("abcd:%s", contents)

	currentURL := url[0:strings.LastIndex(url, "/")]
	for _, m := range matches {

		region := models.Region{
			Name:  string(m[2]),
			Code:  string(m[1]),
			Pcode: "0",
			Type:  "provincetr",
		}

		result.Items = append(result.Items, region)
		result.Requests = append(result.Requests, engine.Request{
			Url:        currentURL + "/" + string(m[1]) + ".html",
			ParserFunc: ParseCityList,
			Pitem:      region,
		})
	}

	return result
}

const cityListRe = `<tr class='([a-z]+)'><td>(<a href='([0-9/]+).html'>)?([0-9]+)(</a>)?</td><td>(<a href='[0-9/]+.html'>)?([^<]+)(</a>)?</td></tr>`

//ParseCityList 处理子元素
func ParseCityList(url string, contents []byte, pitem interface{}) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)

	matches := re.FindAllSubmatch(contents, -1)
	// fmt.Printf("%s\n", contents)

	result := engine.ParseResult{}
	currentURL := url[0:strings.LastIndex(url, "/")]
	pregion := pitem.(models.Region)
	//class m[1]
	//url m[3]
	//code m[4]
	//name m[7]
	// fmt.Printf("%s", matches)
	for _, m := range matches {
		// fmt.Printf("4:%s,7:%s\n", m[4], m[7])
		region := models.Region{
			Code:  string(m[4]),
			Name:  string(m[7]),
			Pcode: pregion.Code,
			Type:  string(m[1]),
		}
		result.Items = append(result.Items, region)

		if m[3] == nil {
			continue
		}

		if string(m[1]) == "towntr" {
			result.Requests = append(result.Requests, engine.Request{
				Url:        currentURL + "/" + string(m[3]) + ".html",
				ParserFunc: ParseVillageList,
				Pitem:      region,
			})
		} else {
			result.Requests = append(result.Requests, engine.Request{
				Url:        currentURL + "/" + string(m[3]) + ".html",
				ParserFunc: ParseCityList,
				Pitem:      region,
			})

		}

	}

	return result
}

const villageListRe = `<tr class='villagetr'><td>([0-9]+)</td><td>[0-9]+</td><td>([^<]+)</td></tr>`

//ParseVillageList 1
func ParseVillageList(url string, contents []byte, pitem interface{}) engine.ParseResult {

	re := regexp.MustCompile(villageListRe)

	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	pregion := pitem.(models.Region)
	for _, m := range matches {
		region := models.Region{
			Code:  string(m[1]),
			Name:  string(m[2]),
			Pcode: pregion.Code,
			Type:  "villagetr",
		}
		result.Items = append(result.Items, region)
	}

	return result
}
