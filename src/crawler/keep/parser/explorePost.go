package parser

import (
	"crawler/engine"
	"crawler/keep/model"
	"encoding/json"
	"regexp"
	"strings"
)

const CONTENT_RE = `var ENTRY = (\{[\w\W]*\});[\w\W]var INAPP_URL`

func ParseExplorePost(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(CONTENT_RE)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		raw := string(m[1])
		raw = strings.TrimPrefix(raw, "{")
		raw = strings.TrimSuffix(raw, "}")
		arr := strings.Split(raw, "\n")
		m := make(map[string]interface{})
		for _, line := range arr {
			if len(line) == 0 {
				continue
			}
			pair := strings.SplitN(line, ":",2)
			k := strings.ReplaceAll(pair[0], " ", "")
			v := strings.TrimSuffix(pair[1], ",")
			v = strings.ReplaceAll(v, " ", "")
			v = strings.TrimSuffix(strings.TrimPrefix(v, "'"), "'")
			m[k] = v
		}
		mjson,_ :=json.Marshal(m)
		var p model.HotPost
		json.Unmarshal(mjson, &p)
		result.Items = append(make([]engine.Item, 0), engine.Item{Id: p.Id, Type: "keep_hot", Url: "", PayLoad: p})
	}
	return result
}

type ExplorePostParser struct {
}

func (p *ExplorePostParser) Parse(contents []byte, url string) engine.ParseResult {
	return ParseExplorePost(contents, "")
}

func (p *ExplorePostParser) Serializ() (name string, args interface{}) {
	return "ParseExplorePost", ""
}

func NewExplorePostParser() *ExplorePostParser {
	return &ExplorePostParser{
	}
}
