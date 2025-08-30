package main

import (
	"log"
	"net/http"

	"spx-integration/config"
	"spx-integration/routes"
)

func main() {
	config.Init()
	r := routes.Init()
	srv := &http.Server{
		Addr:    ":" + config.Cfg.Port,
		Handler: r,
	}
	log.Println("SPX Integration up on :" + config.Cfg.Port)
	log.Fatal(srv.ListenAndServe())
}
