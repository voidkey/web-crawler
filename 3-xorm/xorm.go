package xorm

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	_ "xorm.io/xorm"
)

var engine *xorm.Engine
var err error

type GormPage struct {
	Id      int64
	Title   string
	Content string    `xorm:"text"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

func init() {
	engine, err = xorm.NewEngine("mysql", "root:xzfnk2016@/test_xorm?charset=utf8")
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		err = engine.Ping()
		if err != nil {
			fmt.Printf("err:%v\n", err)
		} else {
			print("Connected!")
		}
	}
}

func fetch(url string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36")
	req.Header.Add("Cookie", "_ga=GA1.1.1322073816.1693711201; _ga_M95P3TTWJZ=GS1.1.1693711200.1.1.1693711262.0.0.0; affinity=1694412311.312.1127.796417|a6728cc07008ec0fd0d6b7ff6028a867; .AspNetCore.Session=CfDJ8Eg9kra6YURKsOjJwROiT4uWU2bQUONAkwRI%2BPzepxn%2BD46FnHsTF7eIk3yDP9jzB%2FuC7YXlxAFD9HPkQUOXAPD6j1UITDFTDsOzi28Cl4w5MgwLBmgneJUv7QvrUAqqBn3eLC951uGTB9A627rWcfjGKNSYDRmwJBXX%2B3Lzd%2FEZ; NotRobot=CfDJ8Eg9kra6YURKsOjJwROiT4sHp7sioprRUh8F3YrYKQ32IRLtdZTOBIMZISP4BFqpqG54FCHmgrlV5sm2mOpO4zxCjl-BoinkuHQ2oksX5reamChfB1Refr2VJY9gqZWOdQ")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("HTTP GET Error:", err)
	}
	if resp.StatusCode != 200 {
		fmt.Println("HTTP status code:", resp.StatusCode)
		return ""
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read Error:", err)
		return ""
	}
	return string(body)
}

func parse(html string) {
	//替换掉空格
	html = strings.Replace(html, "\n", "", -1)
	//边栏内容块正则
	re_sidebar := regexp.MustCompile(`<aside id="sidebar" role="navigation">(.*?)</aside>`)
	//找到边栏块内容
	sidebar := re_sidebar.FindString(html)
	//链接正则
	re_link := regexp.MustCompile(`href="(.*?)"`)
	//找到所有链接
	links := re_link.FindAllString(sidebar, -1)

	base_url := "https://gorm.io/zh_CN/docs/"

	for _, v := range links {
		s := v[6 : len(v)-1]
		url := base_url + s
		fmt.Printf("url: %v\n", url)
	}
}
func parse2(body string) {
	//替换掉空格
	body = strings.Replace(body, "\n", "", -1)
	//页面内容
	re_content := regexp.MustCompile(`<div class="article">(.*?)</div>`)
	//找到页面内容
	content := re_content.FindString(body)
	//标题
	re_title := regexp.MustCompile(`<h1 class="article-title" itemprop="name">(.*?)</h1>`)
	//找到页面标题
	title := re_title.FindString(content)
	fmt.Printf("title: %v\n", title)
	//切片
	title = title[42 : len(title)-5]
	fmt.Printf("title: %v\n", title)
	saveToDB(title, content)
}

func saveToDB(title string, content string) {
	engine.Sync(new(GormPage))

	page := GormPage{
		Title:   title,
		Content: content,
	}
	affected, err := engine.Insert(&page)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	fmt.Println("save:" + string(affected))
}

func Test() {
	url := "https://gorm.io/zh_CN/docs/"
	s := fetch(url)
	parse2(s)
}
