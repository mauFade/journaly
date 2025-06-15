package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	userservice "github.com/mauFade/journaly/internal/application/service/user-service"
	"github.com/mauFade/journaly/internal/presentation/http/handlers"
	"github.com/mauFade/journaly/internal/presentation/http/middleware"
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
	authMiddlware := middleware.NewAuthMiddleware(s.userService)

	userHandler := handlers.NewUserHandler(s.userService)

	s.router.Post("/users", userHandler.CreateUser)
	s.router.Post("/auth", userHandler.Authenticate)

	s.router.Group(func(r chi.Router) {
		r.Use(authMiddlware.Authenticate)

	})
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}
	log.Println("HTTP server running at " + s.port)
	return s.server.ListenAndServe()
}
