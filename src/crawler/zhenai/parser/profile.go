package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)

var nameRe = regexp.MustCompile(`<a class="name fs24">([^<]+)</a>`)
var idRe = regexp.MustCompile(`>ID：([0-9]+)[^<]+<`)
var ageRe = regexp.MustCompile(`<span class="label">年龄：</span>([\d]+)岁`)
var marrigeRe = regexp.MustCompile(`<span class="label">婚况：</span>([^<]+)`)
var heightRe = regexp.MustCompile(`<span class="label">身高：</span>([^<]+)CM`)
var weightRe = regexp.MustCompile(`<span class="label">体重：</span><span field="">([^<]+)KG`)
var incomeRe = regexp.MustCompile(`<span class="label">月收入：</span>([^<]+)`)
var occupationRe = regexp.MustCompile(`<span class="label">职业： </span>([^<]+)<`)
var eduRe = regexp.MustCompile(`<span class="label">学历：</span>([^<]+)`)
var hokouRe = regexp.MustCompile(`<span class="label">籍贯：</span>([^<]+)`)
var xingzuoRe = regexp.MustCompile(`<span class="label">星座：</span><span field="">([^<]+)</span>`)
var houseRe = regexp.MustCompile(`<span class="label">住房条件：</span><span field="">([^<]+)</span>`)
var carRe = regexp.MustCompile(`<span class="label">是否购车：</span><span field="">([^<]+)</span>`)

func ParseProfile(contents []byte, name string, url string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	var age int
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}
	var height int
	height, err = strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}
	var weight int
	weight, err = strconv.Atoi(extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}
	profile.Marrige = extractString(contents, marrigeRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Education = extractString(contents, eduRe)
	profile.Hokou = extractString(contents, hokouRe)
	profile.Xingzuo = extractString(contents, xingzuoRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)
	id := extractString(contents, idRe)

	result := engine.ParseResult{
		Items: []engine.Item{engine.Item{
			Url:     url,
			Id:      id,
			Type:    "zhenai",
			PayLoad: profile,
		}},
	}
	matches := profileRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: NewProfileParser(string(m[2])),
		})
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

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return ParseProfile(contents, url, p.userName)
}

func (p *ProfileParser) Serializ() (name string, args interface{}) {
	return "ParseProfile", p.userName
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		userName: name,
	}
}
