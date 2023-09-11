package fetch

import (
	"fmt"
	"io"
	"net/http"
)

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

func Test() {
	url := "https://zzk.cnblogs.com/s?w=golang"
	str := fetch(url)
	fmt.Println(str)
}
