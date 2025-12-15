package server

import (
	"fmt"
	"net/http"
	"time"
)

func NewHandler(root string, quiet bool) http.Handler {
	fs := http.FileServer(http.Dir(root))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !quiet {
			fmt.Printf("[%s] %s %s\n",
				time.Now().Format("15:04:05"),
				r.Method,
				r.URL.Path)
		}
		secureHeaders(w)
		fs.ServeHTTP(w, r)
	})
}

func secureHeaders(w http.ResponseWriter) {
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "DENY")
}