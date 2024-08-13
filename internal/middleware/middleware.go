package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func LogCall(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Info().Str("Method", r.Method).Str("Path", r.URL.Path).Msg("Got request")
		next.ServeHTTP(w, r)
		log.Info().Str("Method", r.Method).Str("Path", r.URL.Path).Str("Time took", time.Since(start).String()).Msg("Request completed")
	})
}

func GenerateReqId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		iti := uuid.New
		ctx := context.WithValue(r.Context(), "internal-trase-id", iti)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
