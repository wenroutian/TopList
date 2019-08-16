package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"time"
)

func main() {
	//GetZhiHu()
	fmt.Println(GetZhiHu())
}

func GetZhiHu() []map[string]interface{} {
	url := "https://www.v2ex.com/?tab=hot"
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
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + "失败")
		return []map[string]interface{}{}
	}
	var allData []map[string]interface{}
	document.Find(".cell table").Each(func(i int, selection *goquery.Selection) {
		category := selection.Find(".node").Text()
		url, boolUrl := selection.Find(".item_title a").Attr("href")
		text := selection.Find(".item_title a").Text()
		count := selection.Find(".count_livid").Text()

		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": text,
				"url":      "https://www.v2ex.com" + url,
				"count":    count,
				"category": category,
			})
		}
	})
	return allData
}
