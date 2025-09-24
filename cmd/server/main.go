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

	log.Printf("Coucou")

	log.Fatal(http.ListenAndServe(":8080", mux))
}
