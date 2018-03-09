package main

import (
	"crawler/engine"
	"crawler/persistent"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main() {

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueneScheduler{},
		WorkerCount: 100,
		SaverChan:   persistent.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/akesu", //http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}
