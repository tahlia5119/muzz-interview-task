package elastic

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type ElasticClient struct {
	logger *log.Logger
	esAPI  *esapi.API
}

func NewClient(logger *log.Logger, esAPI *esapi.API) *ElasticClient {
	return &ElasticClient{
		logger: logger,
		esAPI:  esAPI,
	}
}

func (es *ElasticClient) performQuery(index string, query map[string]interface{}) error {
	return nil
}
