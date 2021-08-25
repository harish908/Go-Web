package middleware

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.URL)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		// The handler chain will be stopped if your middleware doesn't call next.ServeHTTP() with the corresponding parameters
		// Middlewares should write to ResponseWriter if they are going to terminate the request, and they should not write to ResponseWriter if they are not going to terminate it
		next.ServeHTTP(w, r)
	})
}
