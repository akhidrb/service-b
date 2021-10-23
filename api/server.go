package api

import (
	v1 "khidr/service-b/api/v1"
	"khidr/service-b/service"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server struct {
	listenAddress string
	timeout       time.Duration
	router        *chi.Mux
	service       *service.Service
}

func New(listenAddress string, timeout time.Duration, svc *service.Service) *Server {
	s := &Server{
		listenAddress: listenAddress,
		timeout:       timeout,
		service:       svc,
	}
	s.setupRouter()
	return s
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.listenAddress, s.router)
}

func (s *Server) setupRouter() {
	r := chi.NewRouter()
	r.Use(middleware.Timeout(s.timeout))
	r.Use(middleware.Logger)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	r.Mount("/", v1.Handler(v1.NewV1(s.service)))
	s.router = r
}
