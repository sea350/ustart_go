package main

import (
	"log"
	"os"

	"github.com/sea350/ustart_go/web"
)

var (
	wsPort     = os.Getenv("WS_PORT")
	assetsRoot = os.Getenv("ASSETS_ROOT")
)

func main() {
	log.SetPrefix("Webserver Command, ")
	log.Println("Loading config...")
	cfg := &web.Config{
		Port:       wsPort,
		AssetsRoot: assetsRoot,
	}

	log.Println("Creating new web service from config...")
	ws, err := web.New(cfg)
	if err != nil {
		panic(err)
	}

	log.Println("Running web service...")
	if err = ws.Run(); err != nil {
		log.Printf("Web.Server exited with error: %v\n", err)
		os.Exit(1)
	}
	log.Println("Backend server died peacefully")
}
