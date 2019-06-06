package config

const (
	ItemSaverHost = ":8888"

	ElasticIndex = "keep_explore"

	ItemServerRpc   = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"

	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ParseProfile"
	NilParser     = "NilParser"

	ParseHotPost = "ParseHotPost"
	ParseExplorePost = "ParseExplorePost"
	ParseExplorePostList = "ParseExplorePostList"

	WorkerPort0 = ":9991"
	WorkerPort1 = "9992"
	WorkerPort2 = "9993"
)
