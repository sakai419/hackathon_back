package v1

import (
	"database/sql"
	"fmt"
	"local-test/internal/handlers/account"
	"local-test/internal/handlers/report"
	middleware "local-test/internal/middlewares"
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

func setupAccountRoutes(r *mux.Router, svc *services.Service, client *auth.Client) {
	// Register the account handlers
	h := account.NewAccountHandler(svc)

	// Create options for the account handlers
	opts := account.GorillaServerOptions{
		BaseURL: "",
		BaseRouter: r,
		Middlewares: []account.MiddlewareFunc{
			middleware.AuthMiddleware(client),
		},
	}

	// Register the account handlers
	account.HandlerWithOptions(h, opts)
}

func setupReportRoutes(r *mux.Router, svc *services.Service, client *auth.Client) {
	// Register the report handlers
	h := report.NewReportHandler(svc)

	// Create options for the report handlers
	opts := report.GorillaServerOptions{
		BaseURL: "",
		BaseRouter: r,
		Middlewares: []report.MiddlewareFunc{
			middleware.AuthMiddleware(client),
		},
		ErrorHandlerFunc: report.ErrHandleFunc,
	}

	// Register the report handlers
	report.HandlerWithOptions(h, opts)
}

func SetupRoutes(db *sql.DB, client *auth.Client) *mux.Router {
	r := mux.NewRouter()

	// Create a new router for the /api/v1 path
	apiV1 := r.PathPrefix("/api/v1").Subrouter()
	apiV1.Use(middleware.LoggingMiddleware)

	repo := repositories.NewRepository(db)
	svc := services.NewService(repo)

	// Register the account routes
	setupAccountRoutes(apiV1, svc, client)
	setupReportRoutes(apiV1, svc, client)

	return apiV1
}