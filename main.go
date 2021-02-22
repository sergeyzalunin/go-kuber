// 	Package documentation for API
//	Documentation for go-kuber-api
//
//	Title: The project is based on https://blog.gopheracademy.com/advent-2017/kubernetes-ready-service/
//	Version: unset
//  Schemes: http
//  BasePath: /home
//	Produces:
//		- application/json
//swagger:meta
//go:generate swagger generate spec
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sergeyzalunin/go-kuber/docs"
	"github.com/sergeyzalunin/go-kuber/handlers"
	"github.com/sergeyzalunin/go-kuber/version"
)

func main() {
	docs.SwaggerInfo.Version = version.Release

	log.Printf(
		"Starting the service...\ncommit: %s, build time: %s, release: %s",
		version.Commit, version.BuildTime, version.Release,
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	r := handlers.Router(version.BuildTime, version.Commit, version.Release)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	srv := http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	log.Print("The service is ready to listen and serve.")

	switch <-interrupt {
	case os.Interrupt:
		log.Print("Got SIGINT...")
	case syscall.SIGTERM:
		log.Print("Got SIGTERM...")
	}

	log.Print("The service is shutting down")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Print("Done")
}
