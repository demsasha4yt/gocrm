package gocrm

import (
	"encoding/json"
	"net/http"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
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
		if err := s.service.Users().Create(r.Context(), u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		u.Sanitize()
		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) handleUsersGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.error(w, r, http.StatusNotImplemented, errNoImplemented)
	}
}

func (s *server) handleUsersFind() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.error(w, r, http.StatusNotImplemented, errNoImplemented)
	}
}

func (s *server) handleUsersUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.error(w, r, http.StatusNotImplemented, errNoImplemented)
	}
}

func (s *server) handleUsersDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.error(w, r, http.StatusNotImplemented, errNoImplemented)
	}
}
