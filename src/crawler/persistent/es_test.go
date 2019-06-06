package persistent

import (
	"crawler/engine"
	"crawler/model"
	"testing"

	"gopkg.in/olivere/elastic.v5"
)

func TestSave(t *testing.T) {
	item := engine.Item{
		Url:  "www.baidu.com",
		Id:   "123123",
		Type: "zhenai",
		PayLoad: model.Profile{
			Name:      "紫水晶",   // 姓名
			Gender:    "女",     // 性别 , 可选值: 男, 女
			Age:       32,      // 年龄
			Height:    158,     // 身高， 单位 cm
			Weight:    55,      // 体重， 单位 kg
			Income:    "5001",  // 收入， 单位 元
			Marrige:   "已婚",    // 婚姻状态， 可选值:未婚， 已婚
			Education: "大专",    // 教育程度， 可选值：博士,研究生，本科，高中，初中，小学，幼儿园
			Hokou:     "重庆巴南区", // 住址
		},
	}
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	Save(client, item, "dating_profile")
}
