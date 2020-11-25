package main

import (
	"awesomeProject/server"
	"log"
	"net/http"
	"os"
	"time"
)

const message = "Hello world"

var (
	GoServerAddr = os.Getenv("GO_SERVER_ADDR")
)

func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		t := time.Now()
		writer.Write([]byte(message + " " + t.Format("2006-01-02 15:04:05")))
	})

	srv := server.New(mux, ":8081")
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatalf("server failed to start: %v ", err)
	}
}
