package colly

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func Test() {
	url := "https://gorm.io/zh_CN/docs/"

	c := colly.NewCollector(
	//colly.MaxDepth(1),
	)

	//goquery selector class
	c.OnHTML(".sidebar-link", func(e *colly.HTMLElement) {
		href := e.Attr(("href"))
		if href != "index.html" {
			c.Visit(e.Request.AbsoluteURL(href))
		}
	})

	c.OnHTML(".article-title", func(h *colly.HTMLElement) {
		title := h.Text
		fmt.Println("title:", title)
	})

	c.OnHTML(".article", func(h *colly.HTMLElement) {
		content, _ := h.DOM.Html()
		fmt.Println("content:", content)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(url)

	// // 解析文章url
	// c.OnHTML("h1 a[href]", func(e *colly.HTMLElement) {
	// 	link := e.Attr("href")
	// 	if strings.HasPrefix(e.Request.AbsoluteURL(link), url) {
	// 		text := strings.ReplaceAll(strings.ReplaceAll(e.Text, "\n", ""), " ", "")
	// 		fmt.Printf("Link found: %q -> %s\n", text, e.Request.AbsoluteURL(link))
	// 	}
	// })
	// // 翻页
	// c.OnHTML("nav a[href]", func(e *colly.HTMLElement) {
	// 	next := e.Attr("rel")
	// 	if next == "next" {
	// 		link := e.Attr("href")
	// 		if strings.HasPrefix(e.Request.AbsoluteURL(link), url) {
	// 			fmt.Printf("Page found: %s\n", e.Request.AbsoluteURL(link))
	// 			c.Visit(e.Request.AbsoluteURL(link))
	// 		}
	// 	}
	// })

	c.Visit(url)
}
