package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/wlshcy/rabbit/src/api"
)

var (
	addr = flag.String("addr", "0.0.0.0:8888", "address")
)

func main() {
	flag.Parse()

	srv := &http.Server{
		Addr:           *addr,
		Handler:        api.NewRouter(),
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("[INFO] Server listened on %s", *addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("[ERROR] Server shutdown due to %s", err)
	}
}
