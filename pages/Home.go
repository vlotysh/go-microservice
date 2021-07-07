package pages

import (
	"net/http"
	"time"
)

const message = "Hello world"

func HomePage(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "text/place; charset=utf-8")
	writer.WriteHeader(http.StatusOK)
	t := time.Now()
	writer.Write([]byte(message + " " + t.Format("2006-01-02 15:04:05")))
}