package http

import (
	corehttp "culand-internship/internal/core/http"
	"culand-internship/internal/internship/app"
	"log"
	"net/http"
)

type Handler struct {
	handlers *app.Handlers
}

func NewHttpHandler(handlers *app.Handlers) *Handler {
	return &Handler{handlers: handlers}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	it, err := h.handlers.GetAll.Handle(r.Context(), app.GetAllInternshipsQuery{})
	if err != nil {
		corehttp.RespondError(w, http.StatusInternalServerError, "failed to get internships")
		log.Printf("failed to get internships: %v", err)
		return
	}

	corehttp.RespondSuccess(w, http.StatusOK, it)
}
