package routes

import (
	"lapanxd/webpage-screenshoter/internal/screenshot"
	"log"
	"net/http"
	"runtime/debug"
)

func RegisterScreenshotRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/screenshot", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")
		if url == "" {
			http.Error(w, "Missing url parameter", http.StatusBadRequest)
			return
		}

		img, err := screenshot.TakeScreenshot(url)
		if err != nil {
			log.Printf("TakeScreenshot error: %v\nStack:\n%s", err, debug.Stack())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "image/png")
		w.Write(img)
	})
}
