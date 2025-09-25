package main

import (
	"lapanxd/webpage-screenshoter/internal/routes"
	"log"
	"net/http"

	"github.com/unrolled/secure"
	"golang.org/x/time/rate"
)

func securityHeaders(next http.Handler) http.Handler {
	securityMiddleware := secure.New(secure.Options{
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'",
	})

	return securityMiddleware.Handler(next)
}

var limiter = rate.NewLimiter(5, 10)

func rateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	routes.RegisterScreenshotRoutes(mux)
	routes.RegisterHealthRoutes(mux)

	fileServer := http.FileServer(http.Dir("./public"))
	mux.Handle("/", fileServer)

	handler := securityHeaders(rateLimit(mux))

	log.Printf("Server starting on port %d", 8080)

	log.Fatal(http.ListenAndServe(":8080", handler))

}
