package gocrm

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/demsasha4yt/gocrm.git/internal/app/service"
	"github.com/demsasha4yt/gocrm.git/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
)

type ctxKey int8

const (
	ctxKeyUser ctxKey = iota
	ctxKeyRequestID
	ctxKeyListParams
)

const (
	sessionName = "gocrm_my"
)

var (
	errIncorectEmailOrPassword = errors.New("Неправильный логин или пароль")
	errNotAuthorized           = errors.New("Вы не авторизованы")
	errHasNoRights             = errors.New("Вы не можете этого сделать")
	errNoImplemented           = errors.New("Not implemented")
	errWrongPageParameter      = errors.New("Wrong Page parameter")
	errWrongPageSizeParameter  = errors.New("Wrong PageSize parameter")
)

// Server is a main structure
type server struct {
	logger       *logrus.Logger
	router       *mux.Router
	db           *sql.DB
	store        store.Store
	sessionStore sessions.Store
	service      service.Interface
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func newServer(store store.Store, sessionStore sessions.Store, service service.Interface) *server {
	s := &server{
		logger:       logrus.New(),
		router:       mux.NewRouter(),
		store:        store,
		sessionStore: sessionStore,
		service:      service,
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
			s.WrapRouteMiddlewares(route.Handler(), route.RouteMiddlewares),
		).Methods(route.Method)
	}
	for _, sub := range subrouter.Subrouters {
		s.configureRouter(r, sub)
	}
	s.logger.Infof("%s router loaded", subrouter.Name)
}
