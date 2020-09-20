package gocrm

import (
	"net/http"
	"reflect"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
)

const (
	// AccessCreatePerm ...
	AccessCreatePerm string = "AccessCreate"
	// AccessGetPerm ...
	AccessGetPerm string = "AccessCreate"
	// AccessUpatePerm ....
	AccessUpatePerm string = "AccessUpdate"
	// AccessDeletePerm ...
	AccessDeletePerm string = "AccessDelete"
	// CategoryCreatePerm ...
	CategoryCreatePerm string = "CategoryCreate"
	// CategoryGetPerm ...
	CategoryGetPerm string = "CategoryGet"
	// CategoryUpdatePerm ...
	CategoryUpdatePerm string = "AccessDelete"
	// CategoryDeletePerm ...
	CategoryDeletePerm string = "AccessDelete"
	// CustomerCreatePerm ...
	CustomerCreatePerm string = "AccessDelete"
	// CustomerGetPerm ...
	CustomerGetPerm string = "AccessDelete"
	// CustomerUpdatePerm ...
	CustomerUpdatePerm string = "AccessDelete"
	// CustomerDeletePerm ...
	CustomerDeletePerm string = "AccessDelete"
	// ManufacturerCreatePerm ...
	ManufacturerCreatePerm string = "ManufacturerDelete"
	// ManufacturerGetPerm ...
	ManufacturerGetPerm string = "ManufacturerDelete"
	// ManufacturerUpdatePerm ...
	ManufacturerUpdatePerm string = "ManufacturerDelete"
	// ManufacturerDeletePerm ...
	ManufacturerDeletePerm string = "ManufacturerDelete"
	// OptionCreatePerm ...
	OptionCreatePerm string = "OptionDelete"
	// OptionGetPerm ...
	OptionGetPerm string = "OptionDelete"
	// OptionUpdatePerm ...
	OptionUpdatePerm string = "OptionDelete"
	// OptionDeletePerm ...
	OptionDeletePerm string = "OptionDelete"
	// OrderCreatePerm ...
	OrderCreatePerm string = "OrderDelete"
	// OrderGetPerm ...
	OrderGetPerm string = "OrderDelete"
	// OrderUpdatePerm ...
	OrderUpdatePerm string = "OrderDelete"
	// OrderDeletePerm ...
	OrderDeletePerm string = "OrderDelete"
	// ProductCreatePerm ...
	ProductCreatePerm string = "ProductDelete"
	// ProductGetPerm ...
	ProductGetPerm string = "ProductDelete"
	// ProductUpdatePerm ...
	ProductUpdatePerm string = "ProductDelete"
	// ProductDeletePerm ...
	ProductDeletePerm string = "ProductDelete"
	// UnitCreatePerm ...
	UnitCreatePerm string = "UnitDelete"
	// UnitGetPerm ...
	UnitGetPerm string = "UnitDelete"
	// UnitUpdatePerm ...
	UnitUpdatePerm string = "UnitDelete"
	// UnitDeletePerm ...
	UnitDeletePerm string = "UnitDelete"
	// UserCreatePerm ...
	UserCreatePerm string = "UserDelete"
	// UserGetPerm ...
	UserGetPerm string = "UserDelete"
	// UserUpdatePerm ...
	UserUpdatePerm string = "UserDelete"
	// UserDeletePerm ...
	UserDeletePerm string = "UserDelete"
)

/**
	ACCESS MIDDLEWARES
**/

func (s *server) AccessMiddleware(access ...string) MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, ok := r.Context().Value(ctxKeyUser).(*models.User)
			if !ok {
				s.error(w, r, http.StatusInternalServerError, nil)
				return
			}

			for _, a := range access {
				val := reflect.ValueOf(user.AccessLevel).FieldByName(a)
				if val.Kind() != reflect.Bool {
					s.error(w, r, http.StatusInternalServerError, nil)
					return
				}
				if !val.Bool() {
					s.error(w, r, http.StatusMethodNotAllowed, errHasNoRights)
					return
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}
