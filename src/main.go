package main

import (
	"log"
	"net/http"

	"github.com/wlshcy/rabbit/src/api"
)

func main() {
	srv := &http.Server{
		Addr:           "0.0.0.0:8899",
		Handler:        api.NewRouter(),
		MaxHeaderBytes: 1 << 20,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("[ERROR] Server shutdown due to %s", err)
	}
}
