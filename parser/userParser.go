package parser

import (
	"go_carwler/model"
	"go_carwler/types"
	"regexp"
	"strings"
)

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

func ParseUser(bytes []byte, nameStr string) types.ParseResult {
	user := model.User{}
	user.Name = strings.TrimSpace(nameStr)
	user.Age = extractString(bytes, Age)
	user.Gender = extractString(bytes, Gender)
	user.Marriage = extractString(bytes, Marriage)
	user.Height = extractString(bytes, Height)
	user.Education = extractString(bytes, Education)
	user.Location = extractString(bytes, Location)
	user.Income = extractString(bytes, Income)

	result := types.ParseResult{
		Items: []interface{}{user},
	}
	return result
}

func extractString(bytes []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(bytes)

	if len(match) >= 2 {
		return strings.TrimSpace(string(match[1]))
	} else {
		return ""
	}
}
