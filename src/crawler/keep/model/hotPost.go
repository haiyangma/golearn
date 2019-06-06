package model


type KeepResp struct {
	Ok string       `json:"ok"`
	ErrorCode int32 `json:"errorCode"`
	Text string     `json:"text"`
	Data []HotPost  `json:"data"`
	Version string  `json:"version"`
	LastId string   `json:"lastId"`
}

type Author struct {
	MembershipSchema string `json:"membershipSchema"`
	KeepValue int32 `json:"keepValue"`
	Gender string `json:"gender"`
	VerifiedIconResourceIdWithSide string `json:""`
	AccountType string `json:"accountType"`
	Bio string `json:"bio"`
	MemberStatus int32 `json:"memberStatus"`
	Avatar string `json:"avatar"`
	MaxKeepValue int32 `json:"maxKeepValue"`
	VerifiedIconUrlWithSide string `json:"verifiedIconUrlWithSide"`
	KgLevel int32 `json:"kgLevel"`
	Id string `json:"_id"`
	VerifiedIconResourceId string `json:"verifiedIconResourceId"`
	VerifiedIconUrl string `json:"verifiedIconUrl"`
	WearBadge string `json:"wearBadge"`
	VerifyType int32 `json:"verifyType"`
	Username string `json:"username"`
}

type HotPost struct {
	CardResourceId string `json:"cardResourceId"`
	Country string `json:"country"`
	Reason string `json:"reason"`
	Achievements string `json:"achievements"`
	InternalShareCount int32 `json:"internalShareCount"`
	CheckUnitId string `json:"checkUnitId"`
	StateValue int32 `json:"stateValue"`
	Feel int32 `json:"feel"`
	T string `json:"type"`
	ContentTypeStr string `json:"contentTypeStr"`
	Relation int32 `json:"relation"`
	HashTags string `json:"hashTags"`
	Province string `json:"province"`
	HasFavorited bool `json:"hasFavorited"`
	Id string `json:"id"`
	State string `json:"state"`
	ContentType string `json:"contentType"`
	DayflowBookId string `json:"dayflowBookId"`
	VideoClipThemeId string `json:"videoClipThemeId"`
	Likes int32 `json:"likes"`
	TopicHashtag string `json:"topicHashtag"`
	Images string `json:"images"`
	Author Author `json:"author"`
	Created string `json:"created"`
	SquadId string `json:"squadId"`
	Expansion string `json:"expansion"`
	Quality float64 `json:"quality"`
	Meta string `json:"meta"`
	CommentUsers int32 `json:"commentUsers"`
	District string `json:"district"`
	Vlog bool `json:"vlog"`
	Schema string `json:"schema"`
	City string `json:"city"`
	Timezone string `json:"timezone"`
	SquadTaskId string `json:"squadTaskId"`
	Content string `json:"content"`
	Geo string `json:"geo"`
	RecordResourceId string `json:"recordResourceId"`
	RecordText string `json:"recordText"`
	Public bool `json:"public"`
	Noise bool `json:"noise"`
	SampleComments bool `json:"sampleComments"`
	LikeRank string `json:"likeRank"`
	ViewCount int32 `json:"viewCount"`
	Place string `json:"place"`
	CheckLogId string `json:"checkLogId"`
	IsAd bool `json:"isAd"`
	NationCode string `json:"nationCode"`
	Comments int32 `json:"comments"`
	HasLiked bool `json:"hasLiked"`
	Photo string `json:"photo"`
	SquadDayIndex int32 `json:"squadDayIndex"`
	LikeStatistics string `json:"likeStatistics"`
	ClientIp string `json:"clientIp"`
	Category string `json:"category"`
	ManualViewCount int32 `json:"manualViewCount"`
	ExternalShareCount int32 `json:"externalShareCount"`
	FavoriteCount int32 `json:"favoriteCount"`
}