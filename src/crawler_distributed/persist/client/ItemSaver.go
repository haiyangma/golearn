package client

import (
	"crawler/engine"
	"fmt"
	"log"

	"crawler_distributed/rpcsupport"
	"crawler_distributed/config"
)

func ItemSaver(host string) (itemChan chan engine.Item, err error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	if err != nil {
		return nil, err
	}
	go func() {
		for {
			item := <-out
			fmt.Printf("Got item : %v \n", item)
			var result = ""
			err = client.Call(config.ItemServerRpc, item, &result)
			if err != nil || result != "ok" {
				log.Printf("Item save error saving item %v:%v", item, err)
			}
		}
	}()
	return out, nil
}
