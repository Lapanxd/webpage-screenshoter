package main

import (
	"lapanxd/webpage-screenshoter/internal/routes"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	routes.RegisterScreenshotRoutes(mux)
	routes.RegisterHealthRoutes(mux)

	log.Printf("Server starting on port %d", 8080)

	log.Fatal(http.ListenAndServe(":8080", mux))

}
