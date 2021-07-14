package pages

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

const message = "Hello world"

type Handlers struct {
	logger *log.Logger
}

func (h *Handlers) Home(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-type", "text/place; charset=utf-8")
	writer.WriteHeader(http.StatusOK)

	_, err := writer.Write([]byte(message))

	if err != nil {
		h.logger.Fatalln("Error!")
	}
}

func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("request processed in %s\n", time.Now().Sub(startTime))

		// call route handler
		next(w, r)
	}
}

func (h *Handlers) SetupRoutes(mux *http.ServeMux, r *mux.Router)  {
	r.HandleFunc("/", h.Logger(h.Home)).Methods("GET")
}

func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}