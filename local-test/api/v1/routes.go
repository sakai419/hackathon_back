package v1

import (
	"database/sql"
	"fmt"
	"local-test/internal/account"
	"local-test/internal/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

func NewServer(db *sql.DB) *Server {
	r := SetupRoutes(db)
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

func setupAccountRoutes(r *mux.Router, db *sql.DB) {
	// Create a new account service
	repo := account.NewAccountRepository(db)
	svc := account.NewAccountService(repo)
	handler := account.NewAccountHandler(svc)

	// Create a new router for the /accounts path
	accounts := r.PathPrefix("/accounts").Subrouter()

	// Register the handler functions for the different HTTP methods
	accounts.HandleFunc("", handler.CreateAccount).Methods("POST")
	accounts.HandleFunc("/{account_id}", handler.DeleteAccount).Methods("DELETE")
	accounts.HandleFunc("/{account_id}/suspend", handler.SuspendAccount).Methods("POST")
	accounts.HandleFunc("/{account_id}/unsuspend", handler.UnsuspendAccount).Methods("POST")
}

func SetupRoutes(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	// Create a new router for the /api/v1 path
	apiV1 := r.PathPrefix("/api/v1").Subrouter()
	apiV1.Use(middleware.LoggingMiddleware)

	// Register the account routes
	setupAccountRoutes(apiV1, db)

	return apiV1
}