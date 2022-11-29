package service

import (
	"log"
	"net/http"
	"tifin/securities-search-module/utils/elastic"
)

func SearchHandler(rw http.ResponseWriter, r *http.Request) {
	securityName := r.URL.Query().Get("securityName")

	if len(securityName) == 0 {
		rw.Write([]byte("Security name not provided"))
		return
	}

	esClient := elastic.ElasticClient()

	res, err := esClient.Info()

	if err != nil {
		log.Printf("Failed to get response from client")
		return
	}
	print(res)
	print(securityName, "NAME")
	rw.Write([]byte("Security name provided"))
}
