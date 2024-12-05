package provider

import (
	"digikalajet/utils/elasticsearch"
	"digikalajet/utils/viper"
)

func Provider() {
	viper.InitConfigs()
	elasticsearch.NewElasticSearch()
}
