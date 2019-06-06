package main

import (
	"crawler_distributed/persist"
	"crawler_distributed/rpcsupport"
	"flag"

	"crawler_distributed/config"

	"fmt"

	"gopkg.in/olivere/elastic.v5"
)

var port = flag.Int("port", 0, "the port for itemsaver to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		//fmt.Println("must specify a port")
		//return
		*port = 8888
	}
	serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex)
}

func serveRpc(host, index string) error {
	Client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: Client,
		Index:  index,
	})

}
