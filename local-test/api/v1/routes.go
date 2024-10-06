package v1

import (
	"database/sql"
	"fmt"
	"local-test/internal/handlers"
	"local-test/internal/middlewares"
	"local-test/internal/repositories"
	"local-test/internal/services"
	"log"
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

func NewServer(db *sql.DB, client *auth.Client) *Server {
	r := SetupRoutes(db, client)
	return &Server{router: r}
}

func (s *Server) Router() *mux.Router {
	return s.router
}

func (s *Server) Start(port int) error {
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Starting server on %s", addr)
	return http.ListenAndServe(addr, s.router)
}

func setupAccountRoutes(r *mux.Router, handler *handlers.Handler) {

	// Create a new router for the /accounts path
	accounts := r.PathPrefix("/accounts").Subrouter()

	// Register the handler functions for the different HTTP methods
	accounts.HandleFunc("", handler.CreateAccount).Methods("POST")
	accounts.HandleFunc("/{account_id}", handler.DeleteAccount).Methods("DELETE")
	accounts.HandleFunc("/{account_id}/suspend", handler.SuspendAccount).Methods("POST")
	accounts.HandleFunc("/{account_id}/unsuspend", handler.UnsuspendAccount).Methods("POST")
}

func SetupRoutes(db *sql.DB, client *auth.Client) *mux.Router {
	r := mux.NewRouter()

	// Create a new router for the /api/v1 path
	apiV1 := r.PathPrefix("/api/v1").Subrouter()
	apiV1.Use(middleware.LoggingMiddleware)

	repo := repositories.NewRepository(db)
	svc := services.NewService(repo)
	handler := handlers.NewHandler(svc)

	// Register the account routes
	setupAccountRoutes(apiV1, handler)

	return apiV1
}