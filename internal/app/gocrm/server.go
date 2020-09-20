package gocrm

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/demsasha4yt/gocrm.git/internal/app/store"
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

	s.configureRouter(s.router, s.registerRouters())
	s.logger.Info("Server started.")
	return s
}

func (s *server) configureRouter(router *mux.Router, subrouter *Router) {
	r := router.PathPrefix(subrouter.PathPrefix).Subrouter()
	for _, middleware := range subrouter.middlewares {
		r.Use(mux.MiddlewareFunc(middleware))
	}
	for _, route := range subrouter.Routes {
		s.logger.Infof("%s route loaded, pattern %s", route.Name, route.Pattern)
		r.HandleFunc(
			route.Pattern,
			s.WrapAccessMiddlwares(route.Handler(), route.AccessMiddleware),
		).Methods(route.Method)
	}
	for _, sub := range subrouter.Subrouters {
		s.configureRouter(r, sub)
	}
	s.logger.Infof("%s router loaded", subrouter.Name)
}

// func (s *server) configureRouter() {
// 	// Some middlewares
// 	s.router.Use(s.setRequestID)
// 	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
// 	s.router.Use(s.accessLogMiddleware)
// 	s.router.Use(s.panicMiddleware)

// 	api := s.router.PathPrefix("/api").Subrouter()
// 	api.Use(s.authMiddleware)
// 	// accessRouters := make([]*mux.Router, 0)

// 	// Auth ...
// 	s.router.HandleFunc("/sign_in", s.handleSignIn()).Methods("POST")
// 	s.router.HandleFunc("/logout", s.handleLogout()).Methods("POST")
// 	api.HandleFunc("/info", s.handleWhoAmI()).Methods("GET")

// 	// /api/users/**
// 	api.HandleFunc("/users", s.handleUsersCreate()).Methods("POST", "OPTIONS")
// 	api.HandleFunc("/users", s.handleUsersGet()).Methods("GET", "OPTIONS")
// 	api.HandleFunc("/users/{id:[0-9]+}", s.handleUsersFind()).Methods("GET", "OPTIONS")
// 	api.HandleFunc("/users/{id:[0-9]+}", s.handleUsersUpdate()).Methods("PUT", "OPTIONS")
// 	api.HandleFunc("/users/{id:[0-9]+}", s.handleUsersDelete()).Methods("DELETE", "OPTIONS")

// 	// /api/units/**
// 	api.HandleFunc("/units", s.handleUnitsCreate()).Methods("POST", "OPTIONS")
// 	api.HandleFunc("/units", s.handleUnitsGet()).Methods("GET", "OPTIONS")
// 	api.HandleFunc("/units/{id:[0-9]+}", s.handleUnitsFind()).Methods("GET", "OPTIONS")
// 	api.HandleFunc("/units/{id:[0-9]+}", s.handleUnitsUpdate()).Methods("PUT", "OPTIONS")
// 	api.HandleFunc("/units/{id:[0-9]+}", s.handleUnitsDelete()).Methods("DELETE", "OPTIONS")

// 	// /api/manufacturers
// 	api.HandleFunc("/manufacturers", s.handleManufacturersCreate()).Methods("POST", "OPTIONS")
// 	api.HandleFunc("/manufacturers", s.handleManufacturersGet()).Methods("GET", "OPTIONS").Queries("page", "{page}")
// 	api.HandleFunc("/manufacturers/{id:[0-9]+}", s.handleManufacturersFind()).Methods("GET", "OPTIONS")
// 	api.HandleFunc("/manufacturers/{id:[0-9]+}", s.handleManufacturersUpdate()).Methods("PUT", "OPTIONS")
// 	api.HandleFunc("/manufacturers/{id:[0-9]+}", s.handleManufacturersDelete()).Methods("DELETE", "OPTIONS")

// 	// /api/categorues
// 	api.HandleFunc("/categories", s.handleCategoriesCreate()).Methods("POST", "OPTIONS")
// 	api.HandleFunc("/categories", s.handleCategoriesGet()).Methods("GET", "OPTIONS")
// 	api.HandleFunc("/categories/{id:[0-9]+}", s.handleCategoriesFind()).Methods("GET", "OPTIONS")
// 	api.HandleFunc("/categories/{id:[0-9]+}", s.handleCategoriesUpdate()).Methods("PUT", "OPTIONS")
// 	api.HandleFunc("/categories/{id:[0-9]+}", s.handleCategoriesDelete()).Methods("DELETE", "OPTIONS")

// 	// Serve static files for SPA
// 	spa := &spaHandler{staticPath: "ui/dist", indexPath: "index.html"}
// 	s.router.PathPrefix("/").Handler(spa)

// 	// Check Api Health handler
// 	s.router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
// 		s.respond(w, r, http.StatusOK, map[string]bool{"ok": true})
// 	})
// }
