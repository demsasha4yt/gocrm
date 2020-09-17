package gocrm

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
	"github.com/gorilla/mux"
)

func (s *server) handleManufacturersCreate() http.HandlerFunc {
	type request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Units       []int  `json:"units"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if !s.checkUserAccessRights(r.Context(), models.UserAccessAdmin) {
			s.error(w, r, http.StatusUnauthorized, errHasNoRights)
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
		if err := s.store.Manufacturers().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) handleManufacturersGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !s.checkUserAccessRights(r.Context(), models.UserAccessManager) {
			s.error(w, r, http.StatusUnauthorized, errHasNoRights)
			return
		}
		// TODO: page pagination
		page, err := strconv.Atoi(r.FormValue("page"))
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}
		manufacturers, err := s.store.Manufacturers().FindAll(page)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusOK, manufacturers)
	}
}

func (s *server) handleManufacturersFind() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !s.checkUserAccessRights(r.Context(), models.UserAccessManager) {
			s.error(w, r, http.StatusUnauthorized, errHasNoRights)
			return
		}
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		manufacturer, err := s.store.Manufacturers().Find(id)
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
		if !s.checkUserAccessRights(r.Context(), models.UserAccessAdmin) {
			s.error(w, r, http.StatusUnauthorized, errHasNoRights)
			return
		}
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
		if err := s.store.Manufacturers().Update(id, u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}
		s.respond(w, r, http.StatusOK, u)
	}
}

func (s *server) handleManufacturersDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !s.checkUserAccessRights(r.Context(), models.UserAccessAdmin) {
			s.error(w, r, http.StatusUnauthorized, errHasNoRights)
			return
		}
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		if err := s.store.Manufacturers().Delete(id); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}
		s.respond(w, r, http.StatusOK, "OK")
	}
}
