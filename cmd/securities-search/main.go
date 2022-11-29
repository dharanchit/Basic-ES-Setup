package main

import (
	"fmt"
	"net/http"
	"tifin/securities-search-module/handlers"
	"tifin/securities-search-module/initializers"
)

func init() {
	initializers.LoadConfig()
}

func main() {
	mux := http.NewServeMux()

	handlers.RequestHandler(mux)

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Printf("Error occurred while starting server %s\n", err.Error())
	}

}
