package gocrm

import (
	"context"
	"database/sql"
	"errors"
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

const (
	sessionName = "gocrm_my"
)

var (
	errIncorectEmailOrPassword = errors.New("Неправильный логин или пароль")
	errNotAuthorized           = errors.New("Вы не авторизованы")
	errHasNoRights             = errors.New("Вы не можете этого сделать")
	errNoImplemented           = errors.New("Not implemented")
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

	api := s.router.PathPrefix("/api").Subrouter()
	api.Use(s.authMiddleware)

	// Auth ...
	s.router.HandleFunc("/sign_in", s.handleSignIn()).Methods("POST")
	s.router.HandleFunc("/logout", s.handleLogout()).Methods("POST")
	api.HandleFunc("/info", s.handleWhoAmI()).Methods("GET")

	// /api/users/**
	api.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
	api.HandleFunc("/users", s.handleUsersGet()).Methods("GET")
	api.HandleFunc("/users/{id:[0-9]+}", s.handleUsersFind()).Methods("GET")
	api.HandleFunc("/users/{id:[0-9]+}", s.handleUsersUpdate()).Methods("PUT")
	api.HandleFunc("/users/{id:[0-9]+}", s.handleUsersDelete()).Methods("DELETE")

	// /api/units/**
	api.HandleFunc("/units", s.handleUnitsCreate()).Methods("POST")
	api.HandleFunc("/units", s.handleUnitsGet()).Methods("GET")
	api.HandleFunc("/units/{id:[0-9]+}", s.handleUnitsFind()).Methods("GET")
	api.HandleFunc("/units/{id:[0-9]+}", s.handleUnitsUpdate()).Methods("PUT")
	api.HandleFunc("/units/{id:[0-9]+}", s.handleUnitsDelete()).Methods("DELETE")

	// /api/manufacturers
	api.HandleFunc("/manufacturers", s.handleManufacturersCreate()).Methods("POST")
	api.HandleFunc("/manufacturers", s.handleManufacturersGet()).Methods("GET").Queries("page", "{page}")
	api.HandleFunc("/manufacturers/{id:[0-9]+}", s.handleManufacturersFind()).Methods("GET")
	api.HandleFunc("/manufacturers/{id:[0-9]+}", s.handleManufacturersUpdate()).Methods("PUT")
	api.HandleFunc("/manufacturers/{id:[0-9]+}", s.handleManufacturersDelete()).Methods("DELETE")

	// /api/categorues
	api.HandleFunc("/categories", s.handleCategoriesCreate()).Methods("POST")
	api.HandleFunc("/categories", s.handleCategoriesGet()).Methods("GET")
	api.HandleFunc("/categories/{id:[0-9]+}", s.handleCategoriesFind()).Methods("GET")
	api.HandleFunc("/categories/{id:[0-9]+}", s.handleCategoriesUpdate()).Methods("PUT")
	api.HandleFunc("/categories/{id:[0-9]+}", s.handleCategoriesDelete()).Methods("DELETE")

	// Serve static files for SPA
	spa := &spaHandler{staticPath: "ui/dist", indexPath: "index.html"}
	s.router.PathPrefix("/").Handler(spa)

	// Check Api Health handler
	s.router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, map[string]bool{"ok": true})
	})
}

// checkUserAccessRights checks if user has accessright
func (s *server) checkUserAccessRights(ctx context.Context, accessRight int) bool {
	u, ok := ctx.Value(ctxKeyUser).(*models.User)
	if !ok {
		return false
	}
	return u.HasAccessRight(accessRight)
}
