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
)

func (s *server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
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

func (s *server) handleSessionCreate() http.HandlerFunc {
	type request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u, err := s.store.User().FindByLogin(req.Login)
		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, errIncorectEmailOrPassword)
			return
		}
		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		session.Values["user_id"] = u.ID
		if err := s.sessionStore.Save(r, w, session); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) handleWhoAmI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*models.User))
	}
}