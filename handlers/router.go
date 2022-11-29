package handlers

import (
	"net/http"
	"tifin/securities-search-module/service"
)

func RequestHandler(mux *http.ServeMux) {
	mux.HandleFunc("/v1/search", service.SearchHandler)
}
