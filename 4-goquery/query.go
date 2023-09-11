package goquery

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func Test() {
	url := "https://gorm.io/zh_CN/docs/"
	d, _ := goquery.NewDocument(url)
	d.Find(".sidebar-link").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		base_url := "https://gorm.io/zh_CN/docs/"

		detail_url := base_url + href
		fmt.Println("detail_url:" + detail_url)

		title := d.Find(".article-title").Text()
		//content, _ := d.Find(".article").Html()

		fmt.Println("title:", title)
		//fmt.Println("content:", content)
	})

}
