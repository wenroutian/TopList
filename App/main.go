package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"strings"
	"time"
)

func main() {
	//GetZhiHu()
	fmt.Println(GetZhiHu())
}

func GetZhiHu() []map[string]interface{} {
	url := "https://github.com/trending/php?since=daily"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	res, err := client.Do(request)

	if err != nil {
		fmt.Println("抓取" + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	var allData []map[string]interface{}
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + "失败")
		return []map[string]interface{}{}
	}

	document.Find(".Box article").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find(".lh-condensed a")
		url, boolUrl := s.Attr("href")
		text := s.Text()
		descText := selection.Find("p").Text()

		category := selection.Find(".d-inline-block").AfterHtml("programmingLanguage").Find("span").Text()

		var star, fork string

		selection.Find(".muted-link").Each(func(i int, selection *goquery.Selection) {
			if i == 0 {
				star = selection.Text()
			}
			if i == 1 {
				fork = selection.Text()
			}
		})
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": strings.Trim(text, "\r\n"), "desc":
			descText, "url": "https://github.com" + url,
				"category": category,
				"star":     star,
				"fork":     fork,
			})
		}
	})
	return allData
}
