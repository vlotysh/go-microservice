package main

import (
	"awesomeProject/server"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

const message = "Hello world"

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

var (
	GoServerAddr = os.Getenv("GO_SERVER_ADDR")
)

func HomePage(writer http.ResponseWriter, request *http.Request) {
	t := time.Now()
	writer.Write([]byte(message + " " + t.Format("2006-01-02 15:04:05")))
}

func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomePage)

	srv := server.New(mux, os.Getenv("GO_SERVER_ADDR"))
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatalf("server failed to start: %v ", err)
	}
}

