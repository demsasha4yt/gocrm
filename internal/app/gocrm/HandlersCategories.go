package gocrm

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
	"github.com/demsasha4yt/gocrm.git/internal/app/pagination"
	"github.com/gorilla/mux"
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
		if err := s.service.Categories().Create(r.Context(), u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) handleCategoriesGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pages := pagination.NewFromRequest(r, -1)
		categories, err := s.service.Categories().FindAll(r.Context(), pages.Offset(), pages.Limit())
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusOK, categories)
	}
}

func (s *server) handleCategoriesFind() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		category, err := s.service.Categories().Find(r.Context(), id)
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}
		s.respond(w, r, http.StatusOK, category)
	}
}

func (s *server) handleCategoriesUpdate() http.HandlerFunc {
	type request struct {
		Name        string   `json:"string"`
		Description string   `json:"description"`
		ParentID    null.Int `json:"parent_id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
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
		if err := s.service.Categories().Update(r.Context(), id, u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusOK, u)
	}
}

func (s *server) handleCategoriesDelete() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		if err := s.service.Categories().Delete(r.Context(), id); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusOK, "OK")
	}
}
