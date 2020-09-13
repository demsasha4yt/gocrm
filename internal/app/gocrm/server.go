package gocrm

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
	"github.com/demsasha4yt/gocrm.git/internal/app/store"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
)

type ctxKey int8

const (
	ctxKeyUser ctxKey = iota
	ctxKeyRequestID
)

// Server is a main structure
type server struct {
	logger       *logrus.Logger
	router       *mux.Router
	db           *sql.DB
	store        store.Store
	sessionStore sessions.Store
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func newServer(store store.Store, sessionStore sessions.Store) *server {
	s := &server{
		logger:       logrus.New(),
		router:       mux.NewRouter(),
		store:        store,
		sessionStore: sessionStore,
	}

	s.configureRouter()
	s.logger.Info("Server started.")
	return s
}

func (s *server) configureRouter() {
	// Some middlewares
	s.router.Use(s.setRequestID)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.Use(s.accessLogMiddleware)
	s.router.Use(s.panicMiddleware)

	// Check Api Health handler
	s.router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, map[string]bool{"ok": true})
	})

	// /api/*****
	api := s.router.PathPrefix("/api").Subrouter()
	api.Use(s.authMiddleware)
	api.HandleFunc("/whoami", s.handleWhoAmI()).Methods("GET")
	api.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")

	// Auth ...
	s.router.HandleFunc("/session", s.handleSessionCreate()).Methods("POST")

	// Serve static files for SPA
	spa := &spaHandler{staticPath: "ui/dist", indexPath: "index.html"}
	s.router.PathPrefix("/").Handler(spa)
}

// checkUserAccessRights checks if user has accessright
func (s *server) checkUserAccessRights(ctx context.Context, accessRight int) bool {
	u, ok := ctx.Value(ctxKeyUser).(*models.User)
	if !ok {
		return false
	}
	return u.HasAccessRight(accessRight)
}
