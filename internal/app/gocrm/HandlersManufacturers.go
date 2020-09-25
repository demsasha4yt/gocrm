package gocrm

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
	"github.com/demsasha4yt/gocrm.git/internal/app/pagination"
	"github.com/gorilla/mux"
)

func (s *server) handleManufacturersCreate() http.HandlerFunc {
	type request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Units       []int  `json:"units"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u := &models.Manufacturer{
			Name:        req.Name,
			Description: req.Description,
		}
		for _, unitID := range req.Units {
			u.Units = append(u.Units, &models.Unit{ID: unitID})
		}
		if err := s.service.Manufacturers().Create(r.Context(), u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) handleManufacturersGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pages := pagination.NewFromRequest(r, -1)
		manufacturers, err := s.service.Manufacturers().FindAll(r.Context(), pages.Offset(), pages.Limit())
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusOK, manufacturers)
	}
}

func (s *server) handleManufacturersFind() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		manufacturer, err := s.service.Manufacturers().Find(r.Context(), id)
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}
		s.respond(w, r, http.StatusOK, manufacturer)
	}
}

func (s *server) handleManufacturersUpdate() http.HandlerFunc {
	type request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Units       []int  `json:"units"`
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
		u := &models.Manufacturer{
			Name:        req.Name,
			Description: req.Description,
		}
		for _, unitID := range req.Units {
			u.Units = append(u.Units, &models.Unit{ID: unitID})
		}
		if err := s.service.Manufacturers().Update(r.Context(), id, u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusOK, u)
	}
}

func (s *server) handleManufacturersDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		if err := s.service.Manufacturers().Delete(r.Context(), id); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusOK, "OK")
	}
}
