package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"pxgen.io/user/internal/constants"
	"pxgen.io/user/internal/utils"
	"pxgen.io/user/internal/utils/log"
)

func LogCall(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Infocf(r.Context(), "Request received , PATH = %s , METHOD = %s", r.URL.Path, r.Method)
		next.ServeHTTP(w, r)
		log.Infocf(r.Context(), "Request processed in %s", time.Since(start).String())
	})
}

func GenerateReqId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		iti := uuid.New().String()
		ctx := context.WithValue(r.Context(), constants.TRACE_ID_KEY, iti)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		claims, err := utils.ValidateJWT(authHeader)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
