package v1

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler, adminMW func(http.Handler) http.Handler) {
	mux.HandleFunc("GET /api/v1/internships", h.GetAll)
	mux.HandleFunc("GET /api/v1/internships/{id}", h.GetByID)

	mux.Handle("POST /api/v1/admin/internships", adminMW(http.HandlerFunc(h.Create)))
	mux.Handle("PUT /api/v1/admin/internships/{id}", adminMW(http.HandlerFunc(h.Update)))
	mux.Handle("PATCH /api/v1/admin/internships/{id}", adminMW(http.HandlerFunc(h.Patch)))
}
