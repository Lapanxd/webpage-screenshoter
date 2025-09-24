package routes

import "net/http"

func RegisterHealthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
}
