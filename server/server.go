package server

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/arimatakao/card-validator/validator"
	"github.com/google/uuid"
)

type server struct {
	srv    *http.Server
	logger *slog.Logger
}

func New(address string) server {
	mux := http.NewServeMux()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	mux.HandleFunc("/api/validation", logMiddleware(Validation, logger))

	return server{
		srv: &http.Server{
			Addr:    address,
			Handler: mux,
		},
		logger: logger,
	}
}

func (s server) Run() error {
	return s.srv.ListenAndServe()
}

func (s server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rr *responseRecorder) WriteHeader(statusCode int) {
	rr.statusCode = statusCode
	rr.ResponseWriter.WriteHeader(statusCode)
}

func logMiddleware(next http.HandlerFunc, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		requestId := uuid.New().String()

		logger.Info("Request",
			slog.String("request_id", requestId),
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
		)

		rr := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}

		rr.Header().Set("X-Request-ID", requestId)

		next.ServeHTTP(rr, r)

		logger.Info("Response",
			slog.String("request_id", requestId),
			slog.Int("status", rr.statusCode),
			slog.Duration("duration", time.Since(start)),
		)
	}
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
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func WriteJSON(w http.ResponseWriter, statusCode int, response ApiResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	jsonBody, _ := json.Marshal(response)
	w.Write(jsonBody)
}

func Validation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var card Card

	if err := json.NewDecoder(r.Body).Decode(&card); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	isValid, errValid := validator.IsValid(*card.CardNumber,
		*card.ExpirationMonth, *card.ExpirationYear)

	if errValid != nil {
		WriteJSON(w, http.StatusOK, ApiResponse{
			Valid: isValid,
			Error: &ErrorValidation{
				Code:    errValid.GetCode(),
				Message: errValid.GetMessage(),
			},
		})
		return
	}

	WriteJSON(w, http.StatusOK, ApiResponse{
		Valid: isValid,
		Error: nil,
	})
}
