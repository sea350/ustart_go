package main

import (
	"log"
	"os"

	"github.com/sea350/ustart_go/backend"
)

var (
	backendPort = os.Getenv("BACKEND_PORT")
)

func main() {
	log.SetPrefix("Backend Command, ")
	log.Println("Loading config...")
	cfg := &backend.Config{
		Port: backendPort,
	}

	log.Println("Creating new backend service from config...")
	srv := backend.New(cfg)

	log.Println("Running server...")
	if err := srv.Run(); err != nil {
		log.Printf("Backend server exited with error {%v}\n", err)
		os.Exit(1)
	}

	log.Println("Backend server died peacefully")
}
