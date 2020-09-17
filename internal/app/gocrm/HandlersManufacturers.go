package gocrm

import "net/http"

func (s *server) handleManufacturersCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.error(w, r, http.StatusNotImplemented, errNoImplemented)
	}
}

func (s *server) handleManufacturersGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.error(w, r, http.StatusNotImplemented, errNoImplemented)
	}
}

func (s *server) handleManufacturersFind() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.error(w, r, http.StatusNotImplemented, errNoImplemented)
	}
}

func (s *server) handleManufacturersUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.error(w, r, http.StatusNotImplemented, errNoImplemented)
	}
}

func (s *server) handleManufacturersDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.error(w, r, http.StatusNotImplemented, errNoImplemented)
	}
}
