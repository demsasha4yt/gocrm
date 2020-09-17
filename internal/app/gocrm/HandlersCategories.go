package gocrm

import (
	"encoding/json"
	"net/http"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
	"github.com/guregu/null"
)

func (s *server) handleCategoriesCreate() http.HandlerFunc {
	type request struct {
		Name        string   `json:"string"`
		Description string   `json:"description"`
		ParentID    null.Int `json:"parent_id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u := &models.Category{
			Name:        req.Name,
			Description: req.Description,
			ParentID:    req.ParentID,
		}
		if err := s.store.Categories().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) handleCategoriesGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !s.checkUserAccessRights(r.Context(), models.UserAccessManager) {
			s.error(w, r, http.StatusUnauthorized, errHasNoRights)
			return
		}
		// TODO: page pagination
		categories, err := s.store.Categories().FindAll()
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusOK, categories)
	}
}

func (s *server) handleCategoriesFind() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.error(w, r, http.StatusNotImplemented, errNoImplemented)
	}
}

func (s *server) handleCategoriesUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.error(w, r, http.StatusNotImplemented, errNoImplemented)
	}
}

func (s *server) handleCategoriesDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.error(w, r, http.StatusNotImplemented, errNoImplemented)
	}
}