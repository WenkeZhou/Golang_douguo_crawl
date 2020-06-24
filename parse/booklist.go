package parse

import (
	"../engine"
	"fmt"
	"regexp"
)

// <a class="cookname text-lips " href="/cookbook/185667.html" target="_blank">土豆饼</a>
const CookbookListRe = `<a class="cookname text-lips " href="([^"]+)" target="_blank">([^"]+)</a>`

func ParseCookbookList(content []byte) engine.ParseResult {
	re := regexp.MustCompile(CookbookListRe)

	matches := re.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		fmt.Printf("m0:%v, m1:%v, m2:%v\n", string(m[0]), string(m[1]), string(m[2]))
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url:       "https://www.douguo.com" + string(m[1]),
				ParseFunc: engine.NilParse,
			})
	}

	return result
}
