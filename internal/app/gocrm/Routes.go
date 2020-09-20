package gocrm

import (
	"net/http"

	"github.com/gorilla/handlers"
)

// MiddlewareFunc ...
type MiddlewareFunc func(http.Handler) http.Handler

// MiddlewareFuncs ...
type MiddlewareFuncs []MiddlewareFunc

// Route ...
type Route struct {
	Name             string
	Method           string
	Pattern          string
	AccessMiddleware MiddlewareFuncs
	Handler          func() http.HandlerFunc
}

// Routes ...
type Routes []*Route

// Router ...
type Router struct {
	Name        string
	PathPrefix  string
	middlewares MiddlewareFuncs
	Subrouters  Routers
	Routes      Routes
}

// Routers ...
type Routers []*Router

func (s *server) registerRouters() *Router {
	router := &Router{
		Name:       "Main",
		PathPrefix: "/",
		middlewares: MiddlewareFuncs{
			s.setRequestID,
			handlers.CORS(handlers.AllowedOrigins([]string{"*"})),
			s.accessLogMiddleware,
			s.panicMiddleware,
		},
		Routes: Routes{
			&Route{
				Name:             "Signin",
				Method:           "POST",
				Pattern:          "/signin",
				AccessMiddleware: MiddlewareFuncs{s.AccessMiddleware()},
				Handler:          s.handleSignIn,
			},
			&Route{
				Name:             "Logout",
				Method:           "POST",
				Pattern:          "/logout",
				AccessMiddleware: MiddlewareFuncs{s.AccessMiddleware()},
				Handler:          s.handleLogout,
			},
		},
		Subrouters: Routers{
			&Router{
				Name:        "API",
				PathPrefix:  "/api",
				middlewares: MiddlewareFuncs{s.authMiddleware},
				Routes: Routes{
					// Info about session user
					&Route{
						Name:             "InfoAboutSession",
						Method:           "GET",
						Pattern:          "/info",
						AccessMiddleware: MiddlewareFuncs{s.AccessMiddleware()},
						Handler:          s.handleWhoAmI,
					},
					&Route{
						Name:             "CreateUser",
						Method:           "POST",
						Pattern:          "/users",
						AccessMiddleware: MiddlewareFuncs{s.AccessMiddleware(UserCreatePerm)},
						Handler:          s.handleUsersCreate,
					},
					&Route{
						Name:             "GetAllUsers",
						Method:           "GET",
						Pattern:          "/users",
						AccessMiddleware: MiddlewareFuncs{s.AccessMiddleware(UserGetPerm)},
						Handler:          s.handleUsersGet,
					},
					&Route{
						Name:             "GetAllUsers",
						Method:           "GET",
						Pattern:          "/users/{id:[0-9]+}",
						AccessMiddleware: MiddlewareFuncs{s.AccessMiddleware(UserGetPerm)},
						Handler:          s.handleUsersGet,
					},
					&Route{
						Name:             "UpdateUserByID",
						Method:           "PUT",
						Pattern:          "/users/{id:[0-9]+}",
						AccessMiddleware: MiddlewareFuncs{s.AccessMiddleware(UserUpdatePerm)},
						Handler:          s.handleUsersUpdate,
					},
					&Route{
						Name:             "DeleteUserByID",
						Method:           "DELETE",
						Pattern:          "/users/{id:[0-9]+}",
						AccessMiddleware: MiddlewareFuncs{s.AccessMiddleware(UserDeletePerm)},
						Handler:          s.handleUsersDelete,
					},
				},
			},
		},
	}
	return router
}
