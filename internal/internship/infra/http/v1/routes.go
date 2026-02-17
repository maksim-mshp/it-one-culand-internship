package v1

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("GET /api/v1/internships", h.GetAll)
	mux.HandleFunc("GET /api/v1/internships/{id}", h.GetByID)
	mux.HandleFunc("POST /api/v1/admin/internships", h.Create)
	mux.HandleFunc("PUT /api/v1/admin/internships/{id}", h.Update)
	mux.HandleFunc("PATCH /api/v1/admin/internships/{id}", h.Patch)
}
