package parser

import (
	"crawler/engine"
	"crawler/keep/config"
	"crawler/keep/model"
	"encoding/json"
)

func ParseHotPost(contents []byte, _ string) engine.ParseResult {
	var kr model.KeepResp
	json.Unmarshal(contents,&kr)
	items := make([]engine.Item,0)
	for _,p := range kr.Data {
		item := engine.Item{Id:p.Id,Type:"keep_hot",Url:"",PayLoad:p}
		items = append(items, item)
	}
	r := engine.ParseResult{Items:items}
	if kr.LastId != "" {
		req := engine.Request{Url:config.KEEP_API_HOST+config.KEEP_HOT_URI+"&lastId="+kr.LastId,Parser:engine.NewFuncParser(ParseHotPost, "ParseHotPost")}
		r.Requests = append(make([]engine.Request, 0), req)
	}
	return r
}


type HotPostParser struct {
}

func (p *HotPostParser) Parse(contents []byte, url string) engine.ParseResult {
	return ParseHotPost(contents,"")
}

func (p *HotPostParser) Serializ() (name string, args interface{}) {
	return "ParseHotPost", ""
}

func NewHotPostParser() *HotPostParser {
	return &HotPostParser{
	}
}