package main

import (
	"go-microservice/pages"
	"go-microservice/server"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)


func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main()  {
	logger := log.New(os.Stdout, "go api ", log.LstdFlags | log.Lshortfile)

	h := pages.NewHandlers(logger)

	var (
		GoServerAddr = os.Getenv("GO_SERVER_ADDR")
		GoCertFile   = os.Getenv("GO_CERT_FILE")
		GoKeyFile    = os.Getenv("GO_KEY_FILE")
	)

	mux := http.NewServeMux()

	h.SetupRoutes(mux)

	srv := server.New(mux, GoServerAddr)
	err := srv.ListenAndServeTLS(GoCertFile, GoKeyFile)

	if err != nil {
		log.Fatalf("server failed to start: %v ", err)
	}
}

