package persistent

import (
	"context"

	"crawler/engine"

	"github.com/pkg/errors"
	"gopkg.in/olivere/elastic.v5"
)

func Save(client *elastic.Client, item engine.Item, index string) (result string, err error) {

	if err != nil {
		return "", err
	}

	if item.Type == "" {
		return "", errors.New("must supply Type")
	}
	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	indexService.Do(context.Background())
	if err != nil {
		return "", err
	}
	return "ok", nil
}
