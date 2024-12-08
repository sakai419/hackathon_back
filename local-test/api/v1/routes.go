package v1

import (
	"database/sql"
	"fmt"
	"local-test/internal/config"
	"local-test/internal/handler/account"
	"local-test/internal/handler/admin"
	"local-test/internal/handler/batch"
	"local-test/internal/handler/block"
	"local-test/internal/handler/conversation"
	"local-test/internal/handler/execute"
	"local-test/internal/handler/follow"
	"local-test/internal/handler/notification"
	"local-test/internal/handler/profile"
	"local-test/internal/handler/report"
	"local-test/internal/handler/search"
	"local-test/internal/handler/setting"
	"local-test/internal/handler/sidebar"
	"local-test/internal/handler/tweet"
	"local-test/internal/handler/user"
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
	port   int
}

func NewServer(db *sql.DB, client *auth.Client, cfg *config.ServerConfig) *Server {
	r := SetupRoutes(db, client, cfg.CorsOrigin)
	return &Server{router: r, port: cfg.Port}
}

func (s *Server) Router() *mux.Router {
	return s.router
}

func (s *Server) Start() error {
	addr := fmt.Sprintf(":%d", s.port)
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

func setUpBlockRoutes(r *mux.Router, repo *repository.Repository, svc *service.Service, client *auth.Client) {
	// Register the block handler
	h := block.NewBlockHandler(svc)

	// Create options for the block handler
	opts := block.GorillaServerOptions{
		BaseURL: "",
		BaseRouter: r,
		Middlewares: []block.MiddlewareFunc{
			middleware.AuthClientAndGetInfoMiddleware(repo, client),
			middleware.GetTargetInfoMiddleware(repo),
		},
		ErrorHandlerFunc: block.ErrorHandlerFunc,
	}

	// Register the block handler
	block.HandlerWithOptions(h, opts)
}

func setUpProfileRoutes(r *mux.Router, repo *repository.Repository, svc *service.Service, client *auth.Client) {
	// Register the profile handler
	h := profile.NewProfileHandler(svc)

	// Create options for the profile handler
	opts := profile.GorillaServerOptions{
		BaseURL: "",
		BaseRouter: r,
		Middlewares: []profile.MiddlewareFunc{
			middleware.AuthClientAndGetInfoMiddleware(repo, client),
		},
	}

	// Register the profile handler
	profile.HandlerWithOptions(h, opts)
}

func setUpTweetRoutes(r *mux.Router, repo *repository.Repository, svc *service.Service, client *auth.Client) {
	// Register the tweet handler
	h := tweet.NewTweetHandler(svc)

	// Create options for the tweet handler
	opts := tweet.GorillaServerOptions{
		BaseURL: "",
		BaseRouter: r,
		Middlewares: []tweet.MiddlewareFunc{
			middleware.AuthClientAndGetInfoMiddleware(repo, client),
			middleware.GetTargetInfoMiddleware(repo),
		},
		ErrorHandlerFunc: tweet.ErrorHandlerFunc,
	}

	// Register the tweet handler
	tweet.HandlerWithOptions(h, opts)
}

func setUpUserRoutes(r *mux.Router, repo *repository.Repository, svc *service.Service, client *auth.Client) {
	// Register the user handler
	h := user.NewUserHandler(svc)

	// Create options for the user handler
	opts := user.GorillaServerOptions{
		BaseURL: "",
		BaseRouter: r,
		Middlewares: []user.MiddlewareFunc{
			middleware.AuthClientAndGetInfoMiddleware(repo, client),
			middleware.GetTargetInfoMiddleware(repo),
		},
		ErrorHandlerFunc: user.ErrorHandlerFunc,
	}

	// Register the user handler
	user.HandlerWithOptions(h, opts)
}

func setUpSidebarRoutes(r *mux.Router, repo *repository.Repository, svc *service.Service, client *auth.Client) {
	// Register the sidebar handler
	h := sidebar.NewSidebarHandler(svc)

	// Create options for the sidebar handler
	opts := sidebar.GorillaServerOptions{
		BaseURL: "",
		BaseRouter: r,
		Middlewares: []sidebar.MiddlewareFunc{
			middleware.AuthClientAndGetInfoMiddleware(repo, client),
		},
	}

	// Register the sidebar handler
	sidebar.HandlerWithOptions(h, opts)
}

func setUpSearchRoutes(r *mux.Router, repo *repository.Repository, svc *service.Service, client *auth.Client) {
	// Register the search handler
	h := search.NewSearchHandler(svc)

	// Create options for the search handler
	opts := search.GorillaServerOptions{
		BaseURL: "",
		BaseRouter: r,
		Middlewares: []search.MiddlewareFunc{
			middleware.AuthClientAndGetInfoMiddleware(repo, client),
		},
		ErrorHandlerFunc: search.ErrorHandlerFunc,
	}

	// Register the search handler
	search.HandlerWithOptions(h, opts)
}

func setUpBatchRoutes(r *mux.Router, repo *repository.Repository, svc *service.Service, client *auth.Client) {
	// Register the batch handler
	h := batch.NewBatchHandler(svc)

	// Create options for the batch handler
	opts := batch.GorillaServerOptions{
		BaseURL: "",
		BaseRouter: r,
		Middlewares: []batch.MiddlewareFunc{
			middleware.AuthClientAndGetInfoMiddleware(repo, client),
		},
	}

	// Register the batch handler
	batch.HandlerWithOptions(h, opts)
}

func setUpExecuteRoutes(r *mux.Router, repo *repository.Repository, svc *service.Service, client *auth.Client) {
	// Register the execute handler
	h := execute.NewExecuteHandler(svc)

	// Create options for the execute handler
	opts := execute.GorillaServerOptions{
		BaseURL: "",
		BaseRouter: r,
		Middlewares: []execute.MiddlewareFunc{
			middleware.AuthClientAndGetInfoMiddleware(repo, client),
		},
	}

	// Register the execute handler
	execute.HandlerWithOptions(h, opts)
}

func setUpAdminRoutes(r *mux.Router, repo *repository.Repository, svc *service.Service, client *auth.Client) {
	// Register the admin handler
	h := admin.NewAdminHandler(svc)

	// Create options for the admin handler
	opts := admin.GorillaServerOptions{
		BaseURL: "",
		BaseRouter: r,
		Middlewares: []admin.MiddlewareFunc{
			middleware.AuthClientAndGetInfoMiddleware(repo, client),
			middleware.GetTargetInfoMiddleware(repo),
		},
		ErrorHandlerFunc: admin.ErrorHandlerFunc,
	}

	// Register the admin handler
	admin.HandlerWithOptions(h, opts)
}

func SetupRoutes(db *sql.DB, client *auth.Client, corsOrigin string) *mux.Router {
	r := mux.NewRouter()

	// Create a new router for the /api/v1 path
	apiV1 := r.PathPrefix("/api/v1").Subrouter()
	apiV1.Use(middleware.LoggingMiddleware)
    apiV1.Use(middleware.EnableCorsMiddleware(corsOrigin))

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
	setUpConversationRoutes(apiV1, repo, svc, client)
	setUpBlockRoutes(apiV1, repo, svc, client)
	setUpProfileRoutes(apiV1, repo, svc, client)
	setUpTweetRoutes(apiV1, repo, svc, client)
	setUpUserRoutes(apiV1, repo, svc, client)
	setUpSidebarRoutes(apiV1, repo, svc, client)
	setUpSearchRoutes(apiV1, repo, svc, client)
	setUpBatchRoutes(apiV1, repo, svc, client)
	setUpExecuteRoutes(apiV1, repo, svc, client)
	setUpAdminRoutes(apiV1, repo, svc, client)

	return r
}