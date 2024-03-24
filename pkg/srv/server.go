package srv

import (
	"context"
	"github.com/gorilla/handlers"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	origins := handlers.AllowedOrigins([]string{"*"})
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handlers.CORS(origins)(handler),
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
