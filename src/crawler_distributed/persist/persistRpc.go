package persist

import (
	"crawler/engine"
	"crawler/persistent"

	"fmt"

	"gopkg.in/olivere/elastic.v5"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (service ItemSaverService) Save(item engine.Item, result *string) error {
	r, err := persistent.Save(service.Client, item, service.Index)
	if err != nil {
		return err
	}
	fmt.Printf("succed to save item :%v \n", item)
	*result = r
	return nil
}
