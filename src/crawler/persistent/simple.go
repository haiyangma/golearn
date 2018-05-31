package persistent

import (
	"crawler/engine"
	"fmt"
	"log"

	"gopkg.in/olivere/elastic.v5"
)

func ItemSaver(index string) (itemChan chan engine.Item, err error) {
	out := make(chan engine.Item)
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}
	go func() {
		for {
			item := <-out
			fmt.Printf("Got item : %v \n", item)
			func() {
				_, err = Save(client, item, index)
				if err != nil {
					log.Printf("Item save error saving item %v:%v", item, err)
				}
			}()
		}
	}()
	return out, nil
}
