package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// log.Println("Ami Middleware: aghe print hobo")
			start := time.Now()
			next.ServeHTTP(w, r)
			// end := time.Now()
			// log.Println("Ami Middleware: pore print hobo")

			log.Println("", r.Method, r.URL.Path, time.Since(start))
		},
	)

}
