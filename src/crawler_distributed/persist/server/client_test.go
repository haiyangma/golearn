package main

import (
	"crawler/engine"
	"crawler/model"
	"crawler_distributed/config"
	"crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {

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

	const host = ":8989"
	go serveRpc(host, "dating_test")
	time.Sleep(1 * time.Second)
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	result := ""
	err = client.Call(config.ItemServerRpc,
		item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result:%s;err:%s", result, err)
	}
}
