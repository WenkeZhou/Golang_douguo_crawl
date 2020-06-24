package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func main() {
	//url := "https://www.douban.com/"
	//url := "https://www.douban.com/robots.txt"
	//url := "http://www.google.com/robots.txt"
	//url := "https://book.douban.com/"
	//url := "https://www.douguo.com/cookbook/2509767.html"
	url := "https://www.douguo.com/fenlei"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code: %d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determinEncoding(bodyReader)

	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	result, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s", result)
	parseContent(result)
}

func determinEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("fetch error: %v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func parseContent(content []byte) {
	// <a href="/caipu/水煮肉片" title="水煮肉片" target="_blank">水煮肉片</a>
	//re := regexp.MustCompile(`<a href="([^"]+)" title="([^"]+)" target="_blank">([^"]+)</a>`)
	//str := `<a href="/caipu/水煮肉片" title="水煮肉片" target="_blank">水煮肉片</a>`
	re := regexp.MustCompile(`<a .* href="([^"]+)" title="([^"]+)" target="_blank">([^"]+)</a>`)
	//re := regexp.MustCompile(`href="([^"]+)" title="([^"]+)" target="_blank">`)

	match := re.FindAllSubmatch(content, -1)
	for _, m := range match {
		//fmt.Printf("m[0]:%s m[1]:%s m[2]:%s\n", m[0], m[1], m[2])
		fmt.Printf("%q\n", m)
	}
}
