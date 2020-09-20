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
	CategoryUpdatePerm string = "CategoryUpdate"
	// CategoryDeletePerm ...
	CategoryDeletePerm string = "CategoryDelete"
	// CustomerCreatePerm ...
	CustomerCreatePerm string = "CustomerCreate"
	// CustomerGetPerm ...
	CustomerGetPerm string = "CustomerGet"
	// CustomerUpdatePerm ...
	CustomerUpdatePerm string = "CustomerUpdate"
	// CustomerDeletePerm ...
	CustomerDeletePerm string = "CustomerDelete"
	// ManufacturerCreatePerm ...
	ManufacturerCreatePerm string = "ManufacturerCreate"
	// ManufacturerGetPerm ...
	ManufacturerGetPerm string = "ManufacturerGet"
	// ManufacturerUpdatePerm ...
	ManufacturerUpdatePerm string = "ManufacturerUpdate"
	// ManufacturerDeletePerm ...
	ManufacturerDeletePerm string = "ManufacturerDelete"
	// OptionCreatePerm ...
	OptionCreatePerm string = "OptionCreate"
	// OptionGetPerm ...
	OptionGetPerm string = "OptionGet"
	// OptionUpdatePerm ...
	OptionUpdatePerm string = "OptionUpdate"
	// OptionDeletePerm ...
	OptionDeletePerm string = "OptionDelete"
	// OrderCreatePerm ...
	OrderCreatePerm string = "OrderCreate"
	// OrderGetPerm ...
	OrderGetPerm string = "OrderGet"
	// OrderUpdatePerm ...
	OrderUpdatePerm string = "OrderUpdate"
	// OrderDeletePerm ...
	OrderDeletePerm string = "OrderDelete"
	// ProductCreatePerm ...
	ProductCreatePerm string = "ProductCreate"
	// ProductGetPerm ...
	ProductGetPerm string = "ProductGet"
	// ProductUpdatePerm ...
	ProductUpdatePerm string = "ProductUpdate"
	// ProductDeletePerm ...
	ProductDeletePerm string = "ProductDelete"
	// UnitCreatePerm ...
	UnitCreatePerm string = "UnitCreate"
	// UnitGetPerm ...
	UnitGetPerm string = "UnitGet"
	// UnitUpdatePerm ...
	UnitUpdatePerm string = "UnitUpdate"
	// UnitDeletePerm ...
	UnitDeletePerm string = "UnitDelete"
	// UserCreatePerm ...
	UserCreatePerm string = "UserCreate"
	// UserGetPerm ...
	UserGetPerm string = "UserGet"
	// UserUpdatePerm ...
	UserUpdatePerm string = "UserUpdate"
	// UserDeletePerm ...
	UserDeletePerm string = "UserDelete"
)

/**
	ACCESS MIDDLEWARES
**/

func (s *server) AccessMiddleware(access ...string) AccessMiddleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
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
					s.error(w, r, http.StatusForbidden, errHasNoRights)
					return
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}

func (s *server) WrapAccessMiddlwares(h http.HandlerFunc, m AccessMiddlewares) http.HandlerFunc {
	if len(m) < 1 {
		return h
	}

	wrapped := h

	// loop in reverse to preserve middleware order
	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}

	return wrapped
}
