package main

import (
	"log"
	"cncamp_a04/handler"
	"cncamp_a04/middleware"
	"cncamp_a04/metrics"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/gorilla/mux"
)

func main() {
	log.Print("Starting http server...")

	metrics.Register()

	r := mux.NewRouter()
	r.HandleFunc("/", handler.Tracing)
	r.HandleFunc("/healthz", handler.Healthz)
	r.HandleFunc("/notFound", handler.NotFound)
	r.HandleFunc("/badRequest", handler.BadRequest)
	r.HandleFunc("/mockSignup", handler.Signup).Methods("POST")
	r.Handle("/metrics", promhttp.Handler())
	r.Use(middleware.CORS)
	r.Use(middleware.Delay)
	r.Use(middleware.ResponseHeader)
	r.Use(middleware.Log)

	srv := http.Server{
		Addr: ":8080",
		Handler: r,
	}

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func () {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Print("Server Started")

	<- done
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)

	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shotdown Failed: %+v", err)
	}

	log.Print("Server Exited Properly")
}
