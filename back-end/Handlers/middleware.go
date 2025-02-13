package handlers

import "net/http"

// still need to put the middleware logic to retun the user to login page, idk how it works with the singlepage app

var Middleware = []func(http.HandlerFunc) http.HandlerFunc{
	authMiddleware,
	loginMiddleware,
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Your logic here
		next.ServeHTTP(w, r)
	})
}

func loginMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Your logic here
		next.ServeHTTP(w, r)
	})
}
