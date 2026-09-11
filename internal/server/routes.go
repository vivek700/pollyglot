package server

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/pollyglot/web"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()

	//Routes
	mux.HandleFunc("/", s.HelloWorldHandler)
	mux.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.FS(web.Files))))

	return requestLoggerMiddleware(s.logger)(mux)
}

// func (s *Server) corsMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*") // Replace "*" with specific origins if needed
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
// 		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
// 		w.Header().Set("Access-Control-Allow-Credentials", "false") // Set to "true" if credentials are required
//
// 		// Handle preflight OPTIONS requests
// 		if r.Method == http.MethodOptions {
// 			w.WriteHeader(http.StatusNoContent)
// 			return
// 		}
//
// 		next.ServeHTTP(w, r)
// 	})
//
// }

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{"message": "Hello World"}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResp); err != nil {
		s.logger.Error("Failed to write response", "err", err)
	}
}

func requestLoggerMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info("Request received", "method", r.Method, "path", r.URL.Path)
			next.ServeHTTP(w, r)
		})
	}

}
