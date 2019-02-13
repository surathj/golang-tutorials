package middleware

import (
	"log"
	"net/http"
)

type middleware func(next http.HandlerFunc) http.HandlerFunc

func RunWebServer() {
	lt := chainMiddleware(
		withLogging,
		withTracing,
	)

	http.HandleFunc("/endpoint", lt(Handler))
	log.Println("Running webserver")
	http.ListenAndServe(":8080", nil)
}

func chainMiddleware(mw ...middleware) middleware {

	return func(final http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {
			last := final
			for i := len(mw) - 1; i >= 0; i-- {
				last = mw[i](last)
			}
			last(w, r)
		}
	}
}

func withLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Logging request: ", r.Body)
		next.ServeHTTP(w, r)
	}
}

func withTracing(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Logging tracing")
		next.ServeHTTP(w, r)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello from the handler")
}
