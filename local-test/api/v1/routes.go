package v1

import (
	"database/sql"
	"fmt"
	"local-test/internal/handler/account"
	"local-test/internal/handler/follow"
	"local-test/internal/handler/report"
	"local-test/internal/middleware"
	"local-test/internal/repository"
	"local-test/internal/service"
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

func setUpAccountRoutes(r *mux.Router, svc *service.Service, client *auth.Client) {
	// Register the account handler
	h := account.NewAccountHandler(svc)

	// Create options for the account handler
	opts := account.GorillaServerOptions{
		BaseURL: "",
		BaseRouter: r,
		Middlewares: []account.MiddlewareFunc{
			middleware.AuthMiddleware(client),
		},
	}

	// Register the account handler
	account.HandlerWithOptions(h, opts)
}

func setUpReportRoutes(r *mux.Router, repo *repository.Repository, svc *service.Service, client *auth.Client) {
	// Register the report handler
	h := report.NewReportHandler(svc)

	// Create options for the report handler
	opts := report.GorillaServerOptions{
		BaseURL: "",
		BaseRouter: r,
		Middlewares: []report.MiddlewareFunc{
			middleware.AuthMiddleware(client),
		},
		ErrorHandlerFunc: report.ErrHandleFunc,
	}

	// Register the report handler
	report.HandlerWithOptions(h, opts)
}

func setUpFollowRoutes(r *mux.Router, repo *repository.Repository, svc *service.Service, client *auth.Client) {
	// Register the follow handler
	h := follow.NewFollowHandler(svc)

	// Create options for the follow handler
	opts := follow.GorillaServerOptions{
		BaseURL: "",
		BaseRouter: r,
		Middlewares: []follow.MiddlewareFunc{
			middleware.AuthMiddleware(client),
			middleware.AccountIDMiddleware(repo),
		},
		ErrorHandlerFunc: follow.ErrHandleFunc,
	}

	// Register the follow handler
	follow.HandlerWithOptions(h, opts)
}

func SetupRoutes(db *sql.DB, client *auth.Client) *mux.Router {
	r := mux.NewRouter()

	// Create a new router for the /api/v1 path
	apiV1 := r.PathPrefix("/api/v1").Subrouter()
	apiV1.Use(middleware.LoggingMiddleware)

	repo := repository.NewRepository(db)
	svc := service.NewService(repo)

	// Register the account routes
	setUpAccountRoutes(apiV1, svc, client)
	setUpReportRoutes(apiV1, repo, svc, client)
	setUpFollowRoutes(apiV1, repo, svc, client)

	return apiV1
}