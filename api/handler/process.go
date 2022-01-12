package handler

import (
	"GDGVIT/seo-suggester/pkg/scraper"
	"net/http"
)

// Expose handlers
func MakeUserHandler(r *http.ServeMux) {
	r.Handle("/api/v1/user/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		return
	}))
	r.HandleFunc("/api/v1/search", scraper.MetaTags)

}