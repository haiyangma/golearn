package model


type ExploreResp struct {
	Ok bool `json:"ok"`
	Data ExploreData `json:"data"`
}

type ExploreData struct {
	HasMore bool `json:"hasMore"`
	Html string `json:"html"`
	LastId string `json:"lastId"`
}

type ExplorePost struct {
	Id string `json:"id"`
	Content string `json:"content"`
}
