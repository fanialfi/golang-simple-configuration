package main

import (
	"fmt"
	"log"
	"net/http"
	"simple-comfiguration/conf"
	"time"
)

type CustomMux struct {
	http.ServeMux
}

func (c *CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &c.ServeMux
	if conf.Configuration().Log.Verbose {
		log.Printf(`Incoming request from "%s" accessing "%s"`, r.Host, r.URL.String())
	}

	current.ServeHTTP(w, r)
}

func main() {
	router := new(CustomMux)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World"))
	})
	router.HandleFunc("/howareyou", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("How are you"))
	})

	var handler http.Handler = router
	handler = MiddlewareAuth(handler)

	server := new(http.Server)
	server.Handler = handler
	server.ReadTimeout = conf.Configuration().Server.ReadTimeout * time.Second
	server.WriteTimeout = conf.Configuration().Server.WriteTimeout * time.Second
	server.Addr = fmt.Sprintf(":%d", conf.Configuration().Server.Port)

	if conf.Configuration().Log.Verbose {
		log.Printf("Starting Server at %s\n", server.Addr)
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err.Error())
	}
}
