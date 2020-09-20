package gocrm

import (
	"net/http"

	"github.com/gorilla/handlers"
)

// MiddlewareFunc ...
type MiddlewareFunc func(http.Handler) http.Handler

// MiddlewareFuncs ...
type MiddlewareFuncs []MiddlewareFunc

// AccessMiddleware ...
type AccessMiddleware func(http.HandlerFunc) http.HandlerFunc

// AccessMiddlewares ...
type AccessMiddlewares []AccessMiddleware

// Route ...
type Route struct {
	Name             string
	Method           string
	Pattern          string
	AccessMiddleware AccessMiddlewares
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
				AccessMiddleware: AccessMiddlewares{s.AccessMiddleware()},
				Handler:          s.handleSignIn,
			},
			&Route{
				Name:             "Logout",
				Method:           "POST",
				Pattern:          "/logout",
				AccessMiddleware: AccessMiddlewares{s.AccessMiddleware()},
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
						AccessMiddleware: AccessMiddlewares{s.AccessMiddleware()},
						Handler:          s.handleWhoAmI,
					},
					// User routes
					&Route{
						Name:             "CreateUser",
						Method:           "POST",
						Pattern:          "/users",
						AccessMiddleware: AccessMiddlewares{s.AccessMiddleware(UserCreatePerm)},
						Handler:          s.handleUsersCreate,
					},
					&Route{
						Name:             "GetAllUsers",
						Method:           "GET",
						Pattern:          "/users",
						AccessMiddleware: AccessMiddlewares{s.AccessMiddleware(UserGetPerm)},
						Handler:          s.handleUsersGet,
					},
					&Route{
						Name:             "GetAllUsers",
						Method:           "GET",
						Pattern:          "/users/{id:[0-9]+}",
						AccessMiddleware: AccessMiddlewares{s.AccessMiddleware(UserGetPerm)},
						Handler:          s.handleUsersGet,
					},
					&Route{
						Name:             "UpdateUserByID",
						Method:           "PUT",
						Pattern:          "/users/{id:[0-9]+}",
						AccessMiddleware: AccessMiddlewares{s.AccessMiddleware(UserUpdatePerm)},
						Handler:          s.handleUsersUpdate,
					},
					&Route{
						Name:             "DeleteUserByID",
						Method:           "DELETE",
						Pattern:          "/users/{id:[0-9]+}",
						AccessMiddleware: AccessMiddlewares{s.AccessMiddleware(UserDeletePerm)},
						Handler:          s.handleUsersDelete,
					},
					// Unit Routes
					&Route{
						Name:             "CreateUnit",
						Method:           "POST",
						Pattern:          "/units",
						AccessMiddleware: AccessMiddlewares{s.AccessMiddleware(UnitCreatePerm)},
						Handler:          s.handleUnitsCreate,
					},
					&Route{
						Name:             "GetAllUnit",
						Method:           "GET",
						Pattern:          "/units",
						AccessMiddleware: AccessMiddlewares{s.AccessMiddleware(UnitGetPerm)},
						Handler:          s.handleUnitsGet,
					},
					&Route{
						Name:             "GetUnit",
						Method:           "GET",
						Pattern:          "/users/{id:[0-9]+}",
						AccessMiddleware: AccessMiddlewares{s.AccessMiddleware(UnitGetPerm)},
						Handler:          s.handleUnitsFind,
					},
					&Route{
						Name:             "UpdateUnitByID",
						Method:           "PUT",
						Pattern:          "/users/{id:[0-9]+}",
						AccessMiddleware: AccessMiddlewares{s.AccessMiddleware(UnitUpdatePerm)},
						Handler:          s.handleUnitsUpdate,
					},
					&Route{
						Name:             "DeleteUnitByID",
						Method:           "DELETE",
						Pattern:          "/units/{id:[0-9]+}",
						AccessMiddleware: AccessMiddlewares{s.AccessMiddleware(UnitDeletePerm)},
						Handler:          s.handleUnitsDelete,
					},
					// Manufacrurers Routes
					&Route{
						Name:             "CreateManufacturer",
						Method:           "POST",
						Pattern:          "/manufacturers",
						AccessMiddleware: AccessMiddlewares{s.AccessMiddleware(ManufacturerCreatePerm)},
						Handler:          s.handleManufacturersCreate,
					},
					&Route{
						Name:             "GetAllManufacturers",
						Method:           "GET",
						Pattern:          "/units",
						AccessMiddleware: AccessMiddlewares{s.AccessMiddleware(ManufacturerGetPerm)},
						Handler:          s.handleManufacturersGet,
					},
					&Route{
						Name:             "GetUnit",
						Method:           "GET",
						Pattern:          "/users/{id:[0-9]+}",
						AccessMiddleware: AccessMiddlewares{s.AccessMiddleware(ManufacturerGetPerm)},
						Handler:          s.handleManufacturersFind,
					},
					&Route{
						Name:             "UpdateUnitByID",
						Method:           "PUT",
						Pattern:          "/users/{id:[0-9]+}",
						AccessMiddleware: AccessMiddlewares{s.AccessMiddleware(ManufacturerUpdatePerm)},
						Handler:          s.handleManufacturersUpdate,
					},
					&Route{
						Name:             "DeleteUnitByID",
						Method:           "DELETE",
						Pattern:          "/units/{id:[0-9]+}",
						AccessMiddleware: AccessMiddlewares{s.AccessMiddleware(ManufacturerDeletePerm)},
						Handler:          s.handleManufacturersDelete,
					},
					// Categories Routes

					&Route{
						Name:             "CreateCategory",
						Method:           "POST",
						Pattern:          "/categories",
						AccessMiddleware: AccessMiddlewares{s.AccessMiddleware(ManufacturerCreatePerm)},
						Handler:          s.handleCategoriesCreate,
					},
					&Route{
						Name:             "GetAllCategories",
						Method:           "GET",
						Pattern:          "/categories",
						AccessMiddleware: AccessMiddlewares{s.AccessMiddleware(ManufacturerGetPerm)},
						Handler:          s.handleCategoriesGet,
					},
					&Route{
						Name:             "GetCategory",
						Method:           "GET",
						Pattern:          "/categories/{id:[0-9]+}",
						AccessMiddleware: AccessMiddlewares{s.AccessMiddleware(ManufacturerGetPerm)},
						Handler:          s.handleCategoriesFind,
					},
					&Route{
						Name:             "UpdateCategoryByID",
						Method:           "PUT",
						Pattern:          "/categories/{id:[0-9]+}",
						AccessMiddleware: AccessMiddlewares{s.AccessMiddleware(ManufacturerUpdatePerm)},
						Handler:          s.handleCategoriesUpdate,
					},
					&Route{
						Name:             "DeleteCategoryByID",
						Method:           "DELETE",
						Pattern:          "/categories/{id:[0-9]+}",
						AccessMiddleware: AccessMiddlewares{s.AccessMiddleware(ManufacturerDeletePerm)},
						Handler:          s.handleCategoriesDelete,
					},
				},
			},
		},
	}
	return router
}
