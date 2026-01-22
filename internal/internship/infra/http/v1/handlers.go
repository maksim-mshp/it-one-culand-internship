package v1

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

// @Summary Получить все стажировки
// @Tags internship
// @Accept json
// @Produce json
// @Success 200 {array} InternshipResponse
// @Router /api/v1/internships [get]
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	all, err := h.handlers.GetAll.Handle(r.Context(), app.GetAllInternshipsQuery{})
	if err != nil {
		corehttp.RespondError(w, http.StatusInternalServerError, "failed to get internships")
		log.Printf("failed to get internships: %v", err)
		return
	}

	allDto := make([]InternshipResponse, 0, len(all))
	for _, i := range all {
		allDto = append(allDto, MapInternship(i))
	}

	corehttp.RespondSuccess(w, http.StatusOK, allDto)
}
