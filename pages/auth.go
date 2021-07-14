package pages

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

const message_auth = "Auth route"

type HandlersAuth struct {
	logger *log.Logger
}

func (h *HandlersAuth) Auth(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-type", "text/place; charset=utf-8")
	writer.WriteHeader(http.StatusOK)

	_, err := writer.Write([]byte(message_auth))

	if err != nil {
		h.logger.Fatalln("Error!")
	}
}

func (h *HandlersAuth) AuthId(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	writer.Header().Set("Content-type", "text/place; charset=utf-8")
	writer.WriteHeader(http.StatusOK)

	_, err := writer.Write([]byte(message_auth + " " + vars["id"]))

	if err != nil {
		h.logger.Fatalln("Error!")
	}
}

func (h *HandlersAuth) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("request processed AUTH in %s\n", time.Now().Sub(startTime))

		// call route handler
		next(w, r)
	}
}

func (h *HandlersAuth) SetupRoutes(mux *http.ServeMux, r *mux.Router)  {
	r.HandleFunc("/auth", h.Logger(h.Auth)).Methods("POST")
	r.HandleFunc("/auth/{id}", h.Logger(h.AuthId)).Methods("GET")
}

func NewHandlersAuth(logger *log.Logger) *HandlersAuth {
	return &HandlersAuth{
		logger: logger,
	}
}