package parse

import (
	"../engine"
	"regexp"
)

const regexStr = `<a .* href="([^"]+)" title="([^"]+)" target="_blank">([^"]+)</a>`
const baseUrl = "https://www.douguo.com"

func ParseTag(content []byte) engine.ParseResult {
	// <a href="/caipu/水煮肉片" title="水煮肉片" target="_blank">水煮肉片</a>
	re := regexp.MustCompile(regexStr)

	match := re.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}

	for _, m := range match {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:       baseUrl + string(m[1]),
			ParseFunc: engine.NilParse,
		})
	}

	return result
}
