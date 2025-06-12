package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mauFade/journaly/internal/presentation/http/handlers"
	userservice "github.com/mauFade/journaly/internal/service/user-service"
)

type Server struct {
	router      *chi.Mux
	server      *http.Server
	userService *userservice.UserService
	port        string
}

func NewServer(us *userservice.UserService, p string) *Server {
	return &Server{
		router:      chi.NewRouter(),
		userService: us,
		port:        p,
	}
}

func (s *Server) ConfigureRoutes() {
	userHandler := handlers.NewUserHandler(s.userService)

	s.router.Post("/users", userHandler.CreateUser)
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}
	return s.server.ListenAndServe()
}
