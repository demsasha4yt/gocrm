package gocrm

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
)

const (
	sessionName = "gocrm_my"
)

var (
	errIncorectEmailOrPassword = errors.New("Неправильный логин или пароль")
	errNotAuthorized           = errors.New("Вы не авторизованы")
	errHasNoRights             = errors.New("Вы не можете этого сделать")
)

func (s *server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if !s.checkUserAccessRights(r.Context(), models.UserAccessRRS) {
			s.error(w, r, http.StatusUnauthorized, errHasNoRights)
			return
		}
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u := &models.User{
			Email:    req.Email,
			Login:    req.Login,
			Password: req.Password,
		}
		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		u.Sanitize()
		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) handleUsersUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !s.checkUserAccessRights(r.Context(), models.UserAccessRRS) {
			s.error(w, r, http.StatusUnauthorized, errHasNoRights)
			return
		}
	}
}

func (s *server) handleUserDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !s.checkUserAccessRights(r.Context(), models.UserAccessRRS) {
			s.error(w, r, http.StatusUnauthorized, errHasNoRights)
			return
		}
	}
}
