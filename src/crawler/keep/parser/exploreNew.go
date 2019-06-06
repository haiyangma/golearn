package parser

import (
	"crawler/engine"
	"crawler/keep/config"
	"crawler/keep/model"
	"encoding/json"
	"golang.org/x/net/html"
	"regexp"
	"strings"
)

var liReS = `(<li>[\w\W]*?</li>)`
var idReS = `sensor-item-id=\"(.*)\" sensor-item-reason`
var contentReS = `<div class="text">([\w\W]*?)</div>`
var liRe = regexp.MustCompile(liReS)
var idRe = regexp.MustCompile(idReS)
var contentRe = regexp.MustCompile(contentReS)

func ParseExplorePostListNew(contents []byte, _ string) engine.ParseResult {
	var ep model.ExploreResp
	json.Unmarshal(contents,&ep)
	matches := liRe.FindAllSubmatch([]byte(ep.Data.Html), -1)
	result := engine.ParseResult{}
	for _,m := range matches {
		id := ""
		content := ""
		idMatches := idRe.FindAllSubmatch(m[1],1)
		for _,idm := range idMatches {
			id = string(idm[1])
		}
		contentMatches := contentRe.FindAllSubmatch(m[1],1)
		for _,contentm := range contentMatches {
			content = string(contentm[1])
		}
		content = strings.ReplaceAll(content,"\n","")
		content = strings.ReplaceAll(content,"\r","")
		content = strings.ReplaceAll(content," ","")
		content = strings.ReplaceAll(content,"\t","")
		content = html.EscapeString(content)
		e := model.ExplorePost{id,content}
		result.Items = append(result.Items, engine.Item{id,"explore","",e})
	}
	if ep.Data.LastId != "" {
		result.Requests = append(result.Requests, engine.Request{
			Url:    config.KEEP_HOST+config.KEEP_EXPLORE_URI+"lastId="+ep.Data.LastId,
			Parser: engine.NewFuncParser(ParseExplorePostListNew,"ParseExplorePostListNew"),
		})
	}
	return result
}
