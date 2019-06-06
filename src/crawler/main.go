package main

import (
	"crawler/engine"
	"crawler/keep/config"
	"crawler/persistent"
	"crawler/scheduler"
	"crawler/keep/parser"
)

func main() {
	itemChan, err := persistent.ItemSaver("keep_expore_new")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueneScheduler{},
		WorkerCount:      10,
		SaverChan:        itemChan,
		RequestProcessor: engine.Worker,
	}
	e.Run(engine.Request{
		Url:    config.KEEP_HOST+config.KEEP_EXPLORE_URI, //http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseExplorePostListNew,"ParseExplorePostListNew"),
	})

}
