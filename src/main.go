package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/wlshcy/rabbit/src/api"
	"github.com/wlshcy/rabbit/src/db"
)

var (
	addr  = flag.String("addr", "0.0.0.0:8888", "address")
	mongo = flag.String("mongo", "mongodb://localhost/bluebox", "mongo server addr")
)

func main() {
	flag.Parse()

	mongoDB := db.NewMongoDB(*mongo)
	srv := &http.Server{
		Addr:           *addr,
		Handler:        api.NewRouter(mongoDB),
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("[INFO] Server listened on %s", *addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("[ERROR] Server shutdown due to %s", err)
	}
}
