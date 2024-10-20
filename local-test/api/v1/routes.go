package v1

import (
	"database/sql"
	"fmt"
	"local-test/internal/handler/account"
	"local-test/internal/handler/conversation"
	"local-test/internal/handler/follow"
	"local-test/internal/handler/message"
	"local-test/internal/handler/notification"
	"local-test/internal/handler/report"
	"local-test/internal/handler/setting"
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
			middleware.AuthClientMiddleware(client),
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
			middleware.AuthClientAndGetInfoMiddleware(repo, client),
			middleware.GetTargetInfoMiddleware(repo),
		},
		ErrorHandlerFunc: report.ErrorHandlerFunc,
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
			middleware.AuthClientAndGetInfoMiddleware(repo, client),
			middleware.GetTargetInfoMiddleware(repo),
		},
		ErrorHandlerFunc: follow.ErrorHandlerFunc,
	}

	// Register the follow handler
	follow.HandlerWithOptions(h, opts)
}

func setUpNotificationRoutes(r *mux.Router, repo *repository.Repository, svc *service.Service, client *auth.Client) {
	// Register the notification handler
	h := notification.NewNotificationHandler(svc)

	// Create options for the notification handler
	opts := notification.GorillaServerOptions{
		BaseURL: "",
		BaseRouter: r,
		Middlewares: []notification.MiddlewareFunc{
			middleware.AuthClientAndGetInfoMiddleware(repo, client),
		},
		ErrorHandlerFunc: notification.ErrorHandlerFunc,
	}

	// Register the notification handler
	notification.HandlerWithOptions(h, opts)
}

func setUpSettingRoutes(r *mux.Router, repo *repository.Repository, svc *service.Service, client *auth.Client) {
	// Register the setting handler
	h := setting.NewSettingHandler(svc)

	// Create options for the setting handler
	opts := setting.GorillaServerOptions{
		BaseURL: "",
		BaseRouter: r,
		Middlewares: []setting.MiddlewareFunc{
			middleware.AuthClientAndGetInfoMiddleware(repo, client),
		},
	}

	// Register the setting handler
	setting.HandlerWithOptions(h, opts)
}

func setUpMessageRoutes(r *mux.Router, repo *repository.Repository, svc *service.Service, client *auth.Client) {
	// Register the message handler
	h := message.NewMessageHandler(svc)

	// Create options for the message handler
	opts := message.GorillaServerOptions{
		BaseURL: "",
		BaseRouter: r,
		Middlewares: []message.MiddlewareFunc{
			middleware.AuthClientAndGetInfoMiddleware(repo, client),
			middleware.GetTargetInfoMiddleware(repo),
		},
		ErrorHandlerFunc: message.ErrorHandlerFunc,
	}

	// Register the message handler
	message.HandlerWithOptions(h, opts)
}

func setUpConversationRoutes(r *mux.Router, repo *repository.Repository, svc *service.Service, client *auth.Client) {
	// Register the conversation handler
	h := conversation.NewConversationHandler(svc)

	// Create options for the conversation handler
	opts := conversation.GorillaServerOptions{
		BaseURL: "",
		BaseRouter: r,
		Middlewares: []conversation.MiddlewareFunc{
			middleware.AuthClientAndGetInfoMiddleware(repo, client),
			middleware.GetTargetInfoMiddleware(repo),
		},
		ErrorHandlerFunc: conversation.ErrorHandlerFunc,
	}

	// Register the conversation handler
	conversation.HandlerWithOptions(h, opts)
}

func SetupRoutes(db *sql.DB, client *auth.Client) *mux.Router {
	r := mux.NewRouter()

	// Create a new router for the /api/v1 path
	apiV1 := r.PathPrefix("/api/v1").Subrouter()
	apiV1.Use(middleware.LoggingMiddleware)
    apiV1.Use(middleware.EnableCorsMiddleware)

    // Handle OPTIONS requests
    apiV1.HandleFunc("/{any:.*}", func(w http.ResponseWriter, r *http.Request) {
        http.NotFoundHandler().ServeHTTP(w, r)
    }).Methods(http.MethodOptions)

    // Create a new repository and service
	repo := repository.NewRepository(db)
	svc := service.NewService(repo)

	// Register the account routes
	setUpAccountRoutes(apiV1, svc, client)
	setUpReportRoutes(apiV1, repo, svc, client)
	setUpFollowRoutes(apiV1, repo, svc, client)
	setUpNotificationRoutes(apiV1, repo, svc, client)
	setUpSettingRoutes(apiV1, repo, svc, client)
	setUpMessageRoutes(apiV1, repo, svc, client)
	setUpConversationRoutes(apiV1, repo, svc, client)


	return r
}