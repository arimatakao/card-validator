package server

import (
	"context"
	"encoding/json"
	"net/http"
)

type server struct {
	srv *http.Server
}

func New(address string) server {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/validation", validation)

	return server{
		srv: &http.Server{
			Addr:    address,
			Handler: mux,
		},
	}
}

func (s server) Run() error {
	return s.srv.ListenAndServe()
}

func (s server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

type Card struct {
	CardNumber      *string `json:"card_number,omitempty"`
	ExpirationMonth *string `json:"expiration_month,omitempty"`
	ExpirationYear  *string `json:"expiration_year,omitempty"`
}

type ApiResponse struct {
	Valid bool             `json:"valid"`
	Error *ErrorValidation `json:"error,omitempty"`
}

type ErrorValidation struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func WriteJSON(w http.ResponseWriter, statusCode int, response ApiResponse) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	jsonBody, _ := json.Marshal(response)
	w.Write(jsonBody)
}

func validation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	WriteJSON(w, http.StatusOK, ApiResponse{
		Valid: false,
		Error: &ErrorValidation{
			Code:    123,
			Message: "Test error message",
		},
	})
}
