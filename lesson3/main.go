package main

import (
	"fmt"
	"regexp"
)

func main() {
	//str := "<a class=\"cookname text-lips \" href=\"/cookbook/1588410.html\" target=\"_blank\">虾仁炒西兰花</a>"

	//str := `<table width="690" border="0" cellspacing="0" cellpadding="0" class="retamr br8">
	//<tbody><tr>
	//<td class="lirre" style="border-top:0;">
	//<span class="scname"><a href="/caipu/西兰花" target="_blank">西兰花</a></span>
	//<span class="right scnum">1颗</span>
	//</td>
	//<td style="border-top:0;">
	//<span class="scname"><a href="/caipu/鲜虾" target="_blank">鲜虾</a></span>
	//<span class="right scnum">150g</span>
	//</td>
	//</tr>
	//<tr>
	//<td class="lirre">
	//<span class="scname"><a href="/caipu/胡萝卜" target="_blank">胡萝卜</a></span>
	//<span class="right scnum">半个</span>
	//</td>
	//<td>
	//<span class="scname"><a href="/caipu/料酒" target="_blank">料酒</a></span>
	//<span class="right scnum">适量</span>
	//</td>
	//</tr>
	//<tr>
	//<td class="lirre">
	//<span class="scname"><a href="/caipu/胡椒粉" target="_blank">胡椒粉</a></span>
	//<span class="right scnum">适量</span>
	//</td>
	//<td>
	//<span class="scname"><a href="/caipu/味极鲜酱油" target="_blank">味极鲜酱油</a></span>
	//<span class="right scnum">适量</span>
	//</td>
	//</tr>
	//<tr>
	//<td class="lirre">
	//<span class="scname"><a href="/caipu/盐" target="_blank">盐</a></span>
	//<span class="right scnum">适量</span>
	//</td>
	//<td>
	//<span class="scname"><a href="/caipu/油" target="_blank">油</a></span>
	//<span class="right scnum">适量</span>
	//</td>
	//</tr>
	//<tr>
	//<td class="lirre">
	//<span class="scname"><a href="/caipu/糖" target="_blank">糖</a></span>
	//<span class="right scnum">一小勺</span>
	//</td>
	//<td>
	//<span class="scname"><a href="/caipu/大蒜" target="_blank">大蒜</a></span>
	//<span class="right scnum">2瓣</span>
	//</td>
	//</tr>
	//</tbody></table>`

	strr := `            <table width="690" border="0" cellspacing="0" cellpadding="0" class="retamr br8">
                                                <tr>
                    <td class="lirre" style="border-top:0;">
                        <span class="scname"><a href="/caipu/西兰花" target="_blank">西兰花</a></span>
                        <span class='right scnum'>1颗</span>
                    </td>
                                                                                        <td style="border-top:0;">
                        <span class="scname"><a href="/caipu/鲜虾" target="_blank">鲜虾</a></span>
                        <span class='right scnum'>150g</span>
                    </td>
                </tr>
                                                                                <tr>
                    <td class="lirre" >
                        <span class="scname"><a href="/caipu/胡萝卜" target="_blank">胡萝卜</a></span>
                        <span class='right scnum'>半个</span>
                    </td>
                                                                                        <td >
                        <span class="scname"><a href="/caipu/料酒" target="_blank">料酒</a></span>
                        <span class='right scnum'>适量</span>
                    </td>
                </tr>
                                                                                <tr>
                    <td class="lirre" >
                        <span class="scname"><a href="/caipu/胡椒粉" target="_blank">胡椒粉</a></span>
                        <span class='right scnum'>适量</span>
                    </td>
                                                                                        <td >
                        <span class="scname"><a href="/caipu/味极鲜酱油" target="_blank">味极鲜酱油</a></span>
                        <span class='right scnum'>适量</span>
                    </td>
                </tr>
                                                                                <tr>
                    <td class="lirre" >
                        <span class="scname"><a href="/caipu/盐" target="_blank">盐</a></span>
                        <span class='right scnum'>适量</span>
                    </td>
                                                                                        <td >
                        <span class="scname"><a href="/caipu/油" target="_blank">油</a></span>
                        <span class='right scnum'>适量</span>
                    </td>
                </tr>
                                                                                <tr>
                    <td class="lirre" >
                        <span class="scname"><a href="/caipu/糖" target="_blank">糖</a></span>
                        <span class='right scnum'>一小勺</span>
                    </td>
                                                                                        <td >
                        <span class="scname"><a href="/caipu/大蒜" target="_blank">大蒜</a></span>
                        <span class='right scnum'>2瓣</span>
                    </td>
                </tr>
                                                            </table>
        </div>
                <!-- 步骤 -->
                <div class="step">
            <`

	str2 := `<div class="step">
            <h2 class="mini-title">虾仁炒西兰花的做法</h2>
                        <div class="stepcont clearfix">
                <a class="cboxElement cboxElement2" data-snum="2" rel="recipe_img" href="https://cp1.douguo.com/upload/caiku/b/f/a/yuan_bfc7b5d313c97c246cd9cc222cd3028a.jpeg" data-origin="https://cp1.douguo.com/upload/caiku/b/f/a/yuan_bfc7b5d313c97c246cd9cc222cd3028a.jpeg">
                                            <img class="br8" src="https://cp1.douguo.com/upload/caiku/b/f/a/200_bfc7b5d313c97c246cd9cc222cd3028a.jpeg" alt="虾仁炒西兰花的做法图解1" width="200" height="200">
                                    </a>
                <div class="stepinfo">
                    <p>步骤1</p>
                    将鲜虾洗净，拨壳、去虾线；
                </div>
            </div>
                        <div class="stepcont clearfix">
                <a class="cboxElement cboxElement3" data-snum="3" rel="recipe_img" href="https://cp1.douguo.com/upload/caiku/7/f/8/yuan_7f73d9e06350afc19b3a560bff63a848.jpeg" data-origin="https://cp1.douguo.com/upload/caiku/7/f/8/yuan_7f73d9e06350afc19b3a560bff63a848.jpeg">
                                            <img class="br8" src="https://cp1.douguo.com/upload/caiku/7/f/8/200_7f73d9e06350afc19b3a560bff63a848.jpeg" alt="虾仁炒西兰花的做法图解2" width="200" height="200">
                                    </a>
                <div class="stepinfo">
                    <p>步骤2</p>
                    在拨好的虾仁中加入适量胡椒粉，料酒、味极鲜酱油腌制片刻；
                </div>
            </div>
                        <div class="stepcont clearfix">
                <a class="cboxElement cboxElement4" data-snum="4" rel="recipe_img" href="https://cp1.douguo.com/upload/caiku/6/7/b/yuan_67b01978c2a1e95e4b0058db3e81f38b.jpeg" data-origin="https://cp1.douguo.com/upload/caiku/6/7/b/yuan_67b01978c2a1e95e4b0058db3e81f38b.jpeg">
                                            <img class="br8" src="https://cp1.douguo.com/upload/caiku/6/7/b/200_67b01978c2a1e95e4b0058db3e81f38b.jpeg" alt="虾仁炒西兰花的做法图解3" width="200" height="200">
                                    </a>
                <div class="stepinfo">
                    <p>步骤3</p>
                    将西兰花、胡萝卜洗净，切块；
                </div>
            </div>
                        <div class="stepcont clearfix">
                <a class="cboxElement cboxElement5" data-snum="5" rel="recipe_img" href="https://cp1.douguo.com/upload/caiku/1/2/1/yuan_123d939074fdaaf37cb7551b931956f1.jpeg" data-origin="https://cp1.douguo.com/upload/caiku/1/2/1/yuan_123d939074fdaaf37cb7551b931956f1.jpeg">
                                            <img class="br8" src="https://cp1.douguo.com/upload/caiku/1/2/1/200_123d939074fdaaf37cb7551b931956f1.jpeg" alt="虾仁炒西兰花的做法图解4" width="200" height="200">
                                    </a>
                <div class="stepinfo">
                    <p>步骤4</p>
                    准备一锅清水，放入切好的西兰花，加入适量盐，色拉油，开大火煮开（这样能使西兰花保持翠绿颜色）
                </div>
            </div>
                        <div class="stepcont clearfix">
                <a class="cboxElement cboxElement6" data-snum="6" rel="recipe_img" href="https://cp1.douguo.com/upload/caiku/f/0/6/yuan_f040a3c12172a9e2f3c90030ee0fb876.jpeg" data-origin="https://cp1.douguo.com/upload/caiku/f/0/6/yuan_f040a3c12172a9e2f3c90030ee0fb876.jpeg">
                                            <img class="br8" src="https://cp1.douguo.com/upload/caiku/f/0/6/200_f040a3c12172a9e2f3c90030ee0fb876.jpeg" alt="虾仁炒西兰花的做法图解5" width="200" height="200">
                                    </a>
                <div class="stepinfo">
                    <p>步骤5</p>
                    煮至西兰花变色，捞出沥干水分；
                </div>
            </div>
                        <div class="stepcont clearfix">
                <a class="cboxElement cboxElement7" data-snum="7" rel="recipe_img" href="https://cp1.douguo.com/upload/caiku/d/8/6/yuan_d843152a6f94a4add3ec2b097152fc36.jpeg" data-origin="https://cp1.douguo.com/upload/caiku/d/8/6/yuan_d843152a6f94a4add3ec2b097152fc36.jpeg">
                                            <img class="br8" src="https://cp1.douguo.com/upload/caiku/d/8/6/200_d843152a6f94a4add3ec2b097152fc36.jpeg" alt="虾仁炒西兰花的做法图解6" width="200" height="200">
                                    </a>
                <div class="stepinfo">
                    <p>步骤6</p>
                    热锅倒油，加入大蒜末，腌制好的虾仁，煸炒至金黄色；
                </div>
            </div>
                        <div class="stepcont clearfix">
                <a class="cboxElement cboxElement8" data-snum="8" rel="recipe_img" href="https://cp1.douguo.com/upload/caiku/f/4/8/yuan_f418d094608b140ee4dbd6172833ec58.jpeg" data-origin="https://cp1.douguo.com/upload/caiku/f/4/8/yuan_f418d094608b140ee4dbd6172833ec58.jpeg">
                                            <img class="br8" src="https://cp1.douguo.com/upload/caiku/f/4/8/200_f418d094608b140ee4dbd6172833ec58.jpeg" alt="虾仁炒西兰花的做法图解7" width="200" height="200">
                                    </a>
                <div class="stepinfo">
                    <p>步骤7</p>
                    加入切好的胡萝卜片翻炒一下；
                </div>
            </div>
                        <div class="stepcont clearfix">
                <a class="cboxElement cboxElement9" data-snum="9" rel="recipe_img" href="https://cp1.douguo.com/upload/caiku/f/2/9/yuan_f2287c54bff5f7f4814cf03fd6af9da9.jpeg" data-origin="https://cp1.douguo.com/upload/caiku/f/2/9/yuan_f2287c54bff5f7f4814cf03fd6af9da9.jpeg">
                                            <img class="br8" src="https://cp1.douguo.com/upload/caiku/f/2/9/200_f2287c54bff5f7f4814cf03fd6af9da9.jpeg" alt="虾仁炒西兰花的做法图解8" width="200" height="200">
                                    </a>
                <div class="stepinfo">
                    <p>步骤8</p>
                    加入焯熟的西兰花；
                </div>
            </div>
                        <div class="stepcont clearfix">
                <a class="cboxElement cboxElement10" data-snum="10" rel="recipe_img" href="https://cp1.douguo.com/upload/caiku/4/e/f/yuan_4ea5d43500d89e90826d5b8dde5399ff.jpeg" data-origin="https://cp1.douguo.com/upload/caiku/4/e/f/yuan_4ea5d43500d89e90826d5b8dde5399ff.jpeg">
                                            <img class="br8" src="https://cp1.douguo.com/upload/caiku/4/e/f/200_4ea5d43500d89e90826d5b8dde5399ff.jpeg" alt="虾仁炒西兰花的做法图解9" width="200" height="200">
                                    </a>
                <div class="stepinfo">
                    <p>步骤9</p>
                    加入适量盐、糖，翻炒片刻即可出锅！
                </div>
            </div>
                        <div class="stepcont clearfix">
                <a class="cboxElement cboxElement11" data-snum="11" rel="recipe_img" href="https://cp1.douguo.com/upload/caiku/d/c/1/yuan_dc76328dbd5ea881db874b7eaba37bb1.jpeg" data-origin="https://cp1.douguo.com/upload/caiku/d/c/1/yuan_dc76328dbd5ea881db874b7eaba37bb1.jpeg">
                                            <img class="br8" src="https://cp1.douguo.com/upload/caiku/d/c/1/200_dc76328dbd5ea881db874b7eaba37bb1.jpeg" alt="虾仁炒西兰花的做法图解10" width="200" height="200">
                                    </a>
                <div class="stepinfo">
                    <p>步骤10</p>
                    成品；
                </div>
            </div>
                        <div class="stepcont clearfix">
                <a class="cboxElement cboxElement12" data-snum="12" rel="recipe_img" href="https://cp1.douguo.com/upload/caiku/4/1/e/yuan_416fd6f0de5b70eecbd1d16123d77c9e.jpeg" data-origin="https://cp1.douguo.com/upload/caiku/4/1/e/yuan_416fd6f0de5b70eecbd1d16123d77c9e.jpeg">
                                            <img class="br8" src="https://cp1.douguo.com/upload/caiku/4/1/e/200_416fd6f0de5b70eecbd1d16123d77c9e.jpeg" alt="虾仁炒西兰花的做法图解11" width="200" height="200">
                                    </a>
                <div class="stepinfo">
                    <p>步骤11</p>
                    成品2
                </div>
            </div>
                    </div>`

	str3 := `<div class="author-info left">
                                    <a class="nickname text-lips" href="/u/u20935703262409.html" target="_blank">阅宝麻麻</a>
                                                <a class="proicon" style="position: relative;top: 1px;" href="/user/prodesc" target="_blank">
                    <img class="proimg" src="https://i1.douguo.com/upload/note/d/a/a/da84247847aebe48d9dd0fcdac88862a.png" alt="阅宝麻麻">
                </a>
                            </div>`

	//str1_re := `[\d\D]*?\s*?[\d\D]*?<td.*?>[\d\D]*?<span class="scname"><a href=".*?" target="_blank">([^"]+)</a></span>[\d\D]*?<span class="right scnum">([^"]+)</span>[\d\D]*?</td>`
	str1_re := `<span class="scname"><a href=".*?" target="_blank">([^"]+)</a></span>\s*.*?\s*<span class='right scnum'>([^"]+)</span>`

	//str2_re := `<div class="stepcont clearfix">[\d\D]*?<a.*?class="cboxElement cboxElement.*?".*?href="(.*?)".*?<div class="stepinfo">[\d\D]*?<p>步骤.*?</p>\s*(.*?)\s*</div>`
	//str2_re := `<div class="stepcont clearfix">[\d\D]*?<a.*?class="cboxElement cboxElement.*?".*?href="(.*?)".*?<div class="stepinfo">[\d\D]*?<p>步骤.*?</p>\s*(.*?)\s*</div>`
	str2_re := `<div class="stepcont clearfix">[\d\D]*?<a.*?class="cboxElement cboxElement.*?".*?href="(.*?)".*?[\d\D]*?<div class="stepinfo">[\d\D]*?<p>步骤.*?</p>\s*(.*?)\s*</div>`

	str3_re := `<div class="author-info.*?">[\d\D]*?<a class="nickname text-lips" href="(.*?)" target="_blank">(.*?)</a>\s*.*?\s*<img class="proimg" src="(.*?)"`

	//re := regexp.MustCompile(`<a class="cookname text-lips " href="([^"]+)" target="_blank">([^"]+)</a>`)
	re := regexp.MustCompile(str1_re)

	result := re.FindAllSubmatch([]byte(strr), -1)

	//fmt.Printf("%v\n", result)
	for _, value := range result {
		fmt.Printf("m0:%v\nm1:%v\nm2:%v\n", string(value[0]), string(value[1]), string(value[2]))
		fmt.Println("------------------------------")
	}

	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@")

	re2 := regexp.MustCompile(str2_re)

	result2 := re2.FindAllSubmatch([]byte(str2), -1)

	//fmt.Printf("%v\n", result)
	for _, value := range result2 {
		fmt.Printf("m0:%v\nm1:%v\nm2:%v@@@\n", string(value[0]), string(value[1]), string(value[2]))
		fmt.Println("------------------------------")
	}

	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")

	re3 := regexp.MustCompile(str3_re)

	//result3 := re3.FindAllSubmatch([]byte(str3), -1)
	result3 := re3.FindSubmatch([]byte(str3))

	//fmt.Printf("%v\n", result)
	//for _, value := range result3 {
	//	fmt.Printf("m0:%v\nm1:%v\nm2:%v\nm3:%v@@@\n", string(value), string(value), string(value[2]), string(value))
	//	fmt.Println("------------------------------")
	//}
	fmt.Printf("m0:%v\nm1:%v\nm2:%v\nm3:%v@@@\n", string(result3[0]), string(result3[1]), string(result3[2]), string(result3[3]))
	fmt.Println("------------------------------")
}
