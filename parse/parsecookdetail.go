package parse

import (
	"../engine"
	"../model"
	"fmt"
	"regexp"
)

var authorRe = regexp.MustCompile(`<div class="author-info.*?">[\d\D]*?<a class="nickname text-lips" href="(.*?)" target="_blank">(.*?)</a>\s*.*?\s*<img class="proimg" src="(.*?)"`)

//var foodMaterialRe = regexp.MustCompile(`<td.*?>[\d\D]*?<span class="scname"><a href=".*?" target="_blank">([^"]+)</a></span>[\d\D]*?<span class="right scnum">([^"]+)</span>[\d\D]*?</td>`)
//var foodMaterialRe = regexp.MustCompile(`<td.*?>[\d\D]*?<span class="scname"><a href=".*?" target="_blank">([^"]+)</a></span>[\d\D]*?<span class="right scnum">([^"]+)</span>[\d\D]*?</td>`)
var foodMaterialRe = regexp.MustCompile(`<span class="scname"><a href=".*?" target="_blank">([^"]+)</a></span>\s*.*?\s*<span class='right scnum'>([^"]+)</span>`)
var stepRe = regexp.MustCompile(`<div class="stepcont clearfix">[\d\D]*?<a.*?class="cboxElement cboxElement.*?".*?href="(.*?)".*?[\d\D]*?<div class="stepinfo">[\d\D]*?<p>步骤.*?</p>\s*(.*?)\s*</div>`)

func ParseCookDetail(contents []byte) engine.ParseResult {
	//user := model.User{}
	//cookDetails := model.CookDetail{}
	user := ParseAuthor(contents, authorRe)
	steps := ParseSteps(contents, stepRe)
	foodMaterial := ParseFoodMaterials(contents, foodMaterialRe)

	fmt.Printf("user info:%v \n", user)
	fmt.Printf("cookDetails info:%v \n", steps)
	fmt.Printf("foodMaterial info:%v \n", foodMaterial)

	cookDetail := model.CookDetail{UserInfo: user, StepList: steps, FoodMaterial: foodMaterial}
	result := engine.ParseResult{
		Items: []interface{}{cookDetail},
	}
	return result
}

func ParseAuthor(contents []byte, re *regexp.Regexp) model.User {
	var user model.User
	result := ParseFirstMatch(contents, authorRe)
	if result != nil {
		user.NickName = string(result[1])
		user.HomepageUrl = model.Url("https://www.douguo.com") + model.Url(result[0])
		user.Avatar = model.Url(result[2])
	}
	return user
}

func ParseFoodMaterials(contents []byte, re *regexp.Regexp) []model.Material {
	var material []model.Material
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		material = append(material, model.Material{Material: string(m[1]), Num: string(m[2])})
	}
	return material
}

func ParseSteps(contents []byte, re *regexp.Regexp) []model.Step {
	var steps []model.Step
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		steps = append(steps, model.Step{PicUrl: model.Url(m[1]), Operate: string(m[2])})
	}
	return steps
}

func ParseFirstMatch(contents []byte, re *regexp.Regexp) []string {
	matches := re.FindSubmatch(contents)
	if len(matches) >= 2 {
		var t []string
		for _, m := range matches[1:] {
			t = append(t, string(m))
		}
		return t
	}
	return nil
}

//func ParseAllMatches(contents []byte, re *regexp.Regexp) [][][]byte{
//	matches := re.FindAllSubmatch(contents)
//	if len(matches) > 0 {
//		var t []string
//		for _, m := range matches[1:] {
//			t = append(t, string(m))
//		}
//		return t
//	}
//	return nil
//}
