package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go-microservice/pages"
	"go-microservice/server"
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

	r := mux.NewRouter()

	h := pages.NewHandlers(logger)
	auth := pages.NewHandlersAuth(logger)

	var (
		GoServerAddr = os.Getenv("GO_SERVER_ADDR")
		GoCertFile   = os.Getenv("GO_CERT_FILE")
		GoKeyFile    = os.Getenv("GO_KEY_FILE")
	)

	log.Print("Init web server " + GoServerAddr)

	mux := http.NewServeMux()

	h.SetupRoutes(mux, r)
	auth.SetupRoutes(mux, r)

	mux.Handle("/", r)

	srv := server.New(mux, GoServerAddr)
	err := srv.ListenAndServeTLS(GoCertFile, GoKeyFile)

	if err != nil {
		log.Fatalf("server failed to start: %v ", err)
	}
}

