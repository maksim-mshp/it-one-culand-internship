package v1

import (
	corehttp "culand-internship/internal/core/http"
	"culand-internship/internal/internship/app"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	handlers *app.Handlers
}

func NewHttpHandler(handlers *app.Handlers) *Handler {
	return &Handler{handlers: handlers}
}

// @Summary		Получить все стажировки
// @Tags		internship
// @Success		200 {array} InternshipResponse
// @Router		/internships [get]
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	all, err := h.handlers.GetAll.Handle(r.Context(), app.GetAllInternshipsQuery{})
	if err != nil {
		log.Printf("failed to get internships: %v", err)
		corehttp.RespondError(w, mapError(err))
		return
	}

	allDto := make([]InternshipResponse, 0, len(all))
	for _, i := range all {
		allDto = append(allDto, MapInternship(i))
	}

	corehttp.RespondSuccess(w, http.StatusOK, allDto)
}

// @Summary		Получить стажировку по ID
// @Tags		internship
// @Param		id path int true "ID стажировки"
// @Success		200 {object} InternshipResponse
// @Failure		400 {object} APIError
// @Router		/internships/{id} [get]
func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		corehttp.RespondError(w, corehttp.APIError{
			StatusCode: http.StatusBadRequest,
			Error:      "INVALID_INTERNSHIP_ID",
			Details: map[string]any{
				"id": idStr,
			},
		})
		return
	}

	internship, err := h.handlers.GetByID.Handle(r.Context(), app.GetInternshipByIDQuery{}, id)
	if err != nil {
		log.Printf("failed to get internship: %v", err)
		corehttp.RespondError(w, mapError(err))
		return
	}

	dto := MapInternship(internship)
	corehttp.RespondSuccess(w, http.StatusOK, dto)
}
