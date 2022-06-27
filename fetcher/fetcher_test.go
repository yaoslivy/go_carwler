package fetcher

import (
	"fmt"
	"regexp"
	"testing"
)

/**
<ul class="clearfix user-info">
<li>现居：<span>安徽 合肥</span></li>
<li>籍贯：<span>安徽 合肥</span></li>
<li>星座：<span>巨蟹座</span></li>
<li>生肖：<span>虎</span></li>
<li>身高：<span>181cm</span></li>
<li>血型：<span>回头告诉你</span></li>
<li>职业：<span>回头告诉你</span></li>
<li>收入：<span>2000-5000</span></li>
</ul>
*/
// 预先编译，避免多用户时重复编译
var Age = regexp.MustCompile(`<span class="age s1">(\d+)岁</span>`)
var Gender = regexp.MustCompile(`岁</span>
<span>
([^<]+)</span>
<span class="marrystatus">`)
var Marriage = regexp.MustCompile(`<span class="marrystatus">([^<]+)</span>`)
var Height = regexp.MustCompile(`<span class="height">(\d+)cm</span>`)
var Education = regexp.MustCompile(`<span class="education">([^<]+)</span>`)
var Location = regexp.MustCompile(`现居：<span>([^<]+)</span>`)
var Income = regexp.MustCompile(`收入：<span>([^<]+)</span>`)
var Description = regexp.MustCompile(`<div class="body no-border">
<p>
([^<]+)
</p>
</div>`)

func TestFetch(t *testing.T) {
	//url := "http://album.zhenai.com"
	//url := "http://album.zhenai.com/u/1958903678"
	//url := "https://www.7799520.com/jiaou"
	url := "http://www.7799520.com/user/6053410.html"

	fetch, err := Fetch(url)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(fetch))
	ageage := extractString(fetch, Age)
	GenderGender := extractString(fetch, Gender)
	MarriageMarriage := extractString(fetch, Marriage)
	HeightHeight := extractString(fetch, Height)
	EducationEducation := extractString(fetch, Education)
	LocationLocation := extractString(fetch, Location)
	IncomeIncome := extractString(fetch, Income)
	DescriptionDescription := extractString(fetch, Description)
	fmt.Println("age=", ageage)
	fmt.Println("GenderGender=", GenderGender)
	fmt.Println("MarriageMarriage=", MarriageMarriage)
	fmt.Println("HeightHeight=", HeightHeight)
	fmt.Println("EducationEducation=", EducationEducation)
	fmt.Println("LocationLocation=", LocationLocation)
	fmt.Println("IncomeIncome=", IncomeIncome)
	fmt.Println("DescriptionDescription=", DescriptionDescription)

}

func extractString(bytes []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(bytes)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
