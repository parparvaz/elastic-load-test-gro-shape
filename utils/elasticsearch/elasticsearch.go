package elasticsearch

import (
	"digikalajet/utils/viper"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"sync"
)

type ElasticSearch struct {
	Client *elasticsearch.Client
}

var (
	e    ElasticSearch
	err  error
	once sync.Once
)

func NewElasticSearch() {
	once.Do(func() {
		connect()
	})
}

func connect() {
	config := viper.C().Database.ElasticSearch
	address := fmt.Sprintf("http://%s:%d", config.Address, config.Port)
	e.Client, err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{address},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func E() *ElasticSearch {
	return &e
}
