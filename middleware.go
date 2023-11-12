package main

import "net/http"

const (
	USERNAME = "fanialfi"
	PASSWORD = "saichiopy"
)

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Something went wrong", http.StatusUnauthorized)
			return
		}

		isValid := (username == USERNAME) && (password == PASSWORD)
		if !isValid {
			http.Error(w, "wrong username or password", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
