package parser

import (
	"crawler/engine"
	"crawler/keep/config"
	"crawler/keep/model"
	"encoding/json"
	"regexp"
)

const cityListRe = `"(/entries/.*)" sensor-item-id`

func ParseExplorePostList(contents []byte, _ string) engine.ParseResult {
	var ep model.ExploreResp
	json.Unmarshal(contents,&ep)
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch([]byte(ep.Data.Html), -1)
	result := engine.ParseResult{}
	if ep.Data.LastId != "" {
		result.Requests = append(result.Requests, engine.Request{
			Url:    config.KEEP_HOST+config.KEEP_EXPLORE_URI+"lastId="+ep.Data.LastId,
			Parser: engine.NewFuncParser(ParseExplorePostList,"ParseExplorePostList"),
		})
	}
	for _, m := range matches {
		result.Items = nil //append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:    config.KEEP_HOST+string(m[1]),
			Parser: NewExplorePostParser(),
		})
	}
	return result
}
