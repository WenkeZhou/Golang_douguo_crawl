package fetcher

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
	"time"
)

func DeterminEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("fetch error: %v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

var ratelimit = time.Tick(10 * time.Millisecond)

func Fetch(url string) ([]byte, error) {

	<-ratelimit

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("ERROR: get url:%s", url)
	}
	// user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36")

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code: %d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := DeterminEncoding(bodyReader)

	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

//func Fetch(url string) ([]byte, error) {
//
//	resp, err := http.Get(url)
//	if err != nil {
//		panic(err)
//	}
//
//	defer resp.Body.Close()
//	if resp.StatusCode != http.StatusOK {
//		fmt.Printf("Error status code: %d", resp.StatusCode)
//	}
//
//	bodyReader := bufio.NewReader(resp.Body)
//	e := DeterminEncoding(bodyReader)
//
//	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
//
//	return ioutil.ReadAll(utf8Reader)
//}
//
