package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
)

var nameRe = regexp.MustCompile(`<a class="name fs24">([^<]+)</a>`)
var ageRe = regexp.MustCompile(`<span class="label">年龄：</span>([\d]+)岁`)
var marrigeRe = regexp.MustCompile(`<span class="label">婚况：</span>([^<]+)`)
var heightRe = regexp.MustCompile(`<span class="label">身高：</span>([^<]+)`)
var incomeRe = regexp.MustCompile(`<span class="label">月收入：</span>([^<]+)`)
var occupationRe = regexp.MustCompile(`<span class="label">职业：</span>([^<]+)`)
var eduRe = regexp.MustCompile(`<span class="label">学历：</span>([^<]+)`)
var hokouRe = regexp.MustCompile(`<span class="label">籍贯：</span>([^<]+)`)
var xingzuoRe = regexp.MustCompile(`<span class="label">星座：</span>([^<]+)`)

func paraseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	profile.Marrige = extractString(contents, marrigeRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Education = extractString(contents, eduRe)
	profile.Hokou = extractString(contents, hokouRe)
	profile.Xingzuo = extractString(contents, xingzuoRe)
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
