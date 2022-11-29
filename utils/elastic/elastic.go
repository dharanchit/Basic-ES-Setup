package elastic

import (
	"github.com/elastic/go-elasticsearch/v7"
)

func ElasticClient() *elasticsearch.Client {

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		panic("Failed to establish connection with elastic service")
	}

	return es
}
