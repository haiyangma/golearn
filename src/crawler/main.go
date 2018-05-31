package main

import (
	"crawler/engine"
	"crawler/persistent"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main() {
	itemChan, err := persistent.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueneScheduler{},
		WorkerCount:      100,
		SaverChan:        itemChan,
		RequestProcessor: engine.Worker,
	}
	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun", //http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})

}
