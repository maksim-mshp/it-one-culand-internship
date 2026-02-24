package v1

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler, adminMW, internalMW func(http.Handler) http.Handler) {
	mux.HandleFunc("GET /api/v1/internships", h.GetAll)
	mux.HandleFunc("GET /api/v1/internships/{id}", h.GetByID)

	mux.Handle("GET /api/v1/admin/internships", adminMW(http.HandlerFunc(h.GetAllAdmin)))
	mux.Handle("GET /api/v1/admin/internships/{id}", adminMW(http.HandlerFunc(h.GetByIDAdmin)))
	mux.Handle("POST /api/v1/admin/internships", adminMW(http.HandlerFunc(h.Create)))
	mux.Handle("PUT /api/v1/admin/internships/{id}", adminMW(http.HandlerFunc(h.Update)))
	mux.Handle("PATCH /api/v1/admin/internships/{id}", adminMW(http.HandlerFunc(h.Patch)))
	mux.Handle("PUT /api/v1/admin/internships/{id}/status", adminMW(http.HandlerFunc(h.UpdateStatus)))
	mux.Handle("DELETE /api/v1/admin/internships/{id}", adminMW(http.HandlerFunc(h.Delete)))

	mux.Handle("GET /api/v1/internal/internships/{id}/status", internalMW(http.HandlerFunc(h.GetInternalStatus)))
}
