package worker

import (
	"crawler/engine"
	keepParser "crawler/keep/parser"
	"crawler/zhenai/parser"
	"crawler_distributed/config"

	"fmt"

	"log"

	"github.com/pkg/errors"
)

type SerializedParser struct {
	FunctionName string
	Args         interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serializ()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			FunctionName: name,
			Args:         args,
		},
	}
}
func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(request Request) (engine.Request, error) {
	parserFunc, err := DeserializeParser(&request.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    request.Url,
		Parser: parserFunc,
	}, nil
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		enginqRequest, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error desialize request %v:", req)
			continue
		}
		result.Requests = append(result.Requests, enginqRequest)
	}
	return result
}

func DeserializeParser(p *SerializedParser) (engine.Parser, error) {
	switch p.FunctionName {
	case config.ParseCity:
		return engine.NewFuncParser(
			parser.ParseCity,
			config.ParseCity), nil
	case config.ParseCityList:
		return engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList), nil
	case config.ParseProfile:
		if userName, ok := p.Args.(string); ok {
			return parser.NewProfileParser(userName), nil
		} else {
			return nil, fmt.Errorf("invalid arg %v:", p.Args)
		}
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseHotPost:
		return keepParser.NewHotPostParser(),nil
	case config.ParseExplorePost:
		return keepParser.NewExplorePostParser(),nil
	case config.ParseExplorePostList:
		return engine.NewFuncParser(
			keepParser.ParseExplorePostList,
			config.ParseExplorePostList), nil
	default:
		return nil, errors.New("unknow parser name")
	}
}

