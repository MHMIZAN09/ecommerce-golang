package global_router

import "net/http"

func GlobalRouter(mux *http.ServeMux) http.Handler {
	handlerAllReq := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			return
		} else {
			mux.ServeHTTP(w, r)
		}

	}
	return http.HandlerFunc(handlerAllReq)
}

// func CorsMiddleware(next http.Handler) http.Handler {
// 	handler := func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 		w.Header().Set("Content-Type", "application/json")
// 		next.ServeHTTP(w, r)
// 	}
// 	return http.HandlerFunc(handler)
// }
