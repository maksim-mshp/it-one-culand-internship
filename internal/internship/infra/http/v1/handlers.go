package v1

import (
	corehttp "culand-internship/internal/core/http"
	"culand-internship/internal/internship/app"
	"culand-internship/internal/internship/app/handlers"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	handlers *handlers.Handlers
}

func NewHttpHandler(handlers *handlers.Handlers) *Handler {
	return &Handler{handlers: handlers}
}

// @Summary		Получить все стажировки
// @Tags		internship
// @Success		200 {array} InternshipDto
// @Router		/internships [GET]
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	all, err := h.handlers.GetAll.Handle(r.Context(), app.GetAllInternshipsQuery{})
	if err != nil {
		log.Printf("failed to get internships: %v", err)
		corehttp.RespondError(w, mapError(err))
		return
	}

	allDto := make([]InternshipDto, 0, len(all))
	for _, i := range all {
		allDto = append(allDto, MapInternship(i))
	}

	corehttp.Respond(w, http.StatusOK, allDto)
}

// @Summary		Получить стажировку по ID
// @Tags		internship
// @Param		id path int true "ID стажировки"
// @Success		200 {object} InternshipDto
// @Failure		400 {object} APIError
// @Failure		404 {object} APIError
// @Router		/internships/{id} [GET]
func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		corehttp.RespondError(w, corehttp.APIError{
			StatusCode: http.StatusBadRequest,
			Error:      "INVALID_INTERNSHIP_ID",
			Details: map[string]any{
				"id": idStr,
			},
		})
		return
	}

	internship, err := h.handlers.GetByID.Handle(r.Context(), app.GetInternshipByIDQuery{ID: id})
	if err != nil {
		log.Printf("failed to get internship: %v", err)
		corehttp.RespondError(w, mapError(err))
		return
	}

	dto := MapInternship(internship)
	corehttp.Respond(w, http.StatusOK, dto)
}

// @Summary		Создать стажировку
// @Tags		internship
// @Param		internship body InternshipRequestDto true "Стажировка"
// @Success		201 {object} InternshipDto
// @Failure		401 {object} APIError
// @Failure		422 {object} APIError
// @Router		/admin/internships [POST]
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var internshipDto InternshipRequestDto
	apiErr := corehttp.ParseJSONBody(r, &internshipDto)
	if apiErr != nil {
		corehttp.RespondError(w, *apiErr)
		return
	}

	internship, err := h.handlers.Create.Handle(r.Context(), app.CreateInternshipCommand{
		Title:       internshipDto.Title,
		Label:       internshipDto.Label,
		Description: internshipDto.Description,
		Skills:      internshipDto.Skills,
		Goals:       internshipDto.Goals,
	})
	if err != nil {
		log.Printf("failed to create internship: %v", err)
		corehttp.RespondError(w, mapError(err))
		return
	}

	dto := MapInternship(internship)
	location := fmt.Sprintf("/api/v1/internships/%d", internship.ID())
	w.Header().Set("Location", location)
	corehttp.Respond(w, http.StatusCreated, dto)
}

// @Summary		Обновить стажировку
// @Tags		internship
// @Param		id path int true "ID стажировки"
// @Param		internship body InternshipRequestDto true "Стажировка"
// @Success		200 {object} InternshipDto
// @Failure		400 {object} APIError
// @Failure		401 {object} APIError
// @Failure		404 {object} APIError
// @Failure		422 {object} APIError
// @Router		/admin/internships/{id} [PUT]
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		corehttp.RespondError(w, corehttp.APIError{
			StatusCode: http.StatusBadRequest,
			Error:      "INVALID_INTERNSHIP_ID",
			Details: map[string]any{
				"id": idStr,
			},
		})
		return
	}

	var internshipDto InternshipRequestDto
	apiErr := corehttp.ParseJSONBody(r, &internshipDto)
	if apiErr != nil {
		corehttp.RespondError(w, *apiErr)
		return
	}

	internship, err := h.handlers.Update.Handle(r.Context(), app.UpdateInternshipCommand{
		ID:          id,
		Title:       internshipDto.Title,
		Label:       internshipDto.Label,
		Description: internshipDto.Description,
		Skills:      internshipDto.Skills,
		Goals:       internshipDto.Goals,
	})

	if err != nil {
		log.Printf("failed to update internship: %v", err)
		corehttp.RespondError(w, mapError(err))
		return
	}

	dto := MapInternship(internship)
	corehttp.Respond(w, http.StatusOK, dto)
}

// @Summary		Обновить стажировку
// @Tags		internship
// @Param		id path int true "ID стажировки"
// @Param		internship body InternshipRequestDto true "Стажировка"
// @Success		200 {object} InternshipDto
// @Failure		400 {object} APIError
// @Failure		401 {object} APIError
// @Failure		404 {object} APIError
// @Failure		422 {object} APIError
// @Router		/admin/internships/{id} [PATCH]
func (h *Handler) Patch(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		corehttp.RespondError(w, corehttp.APIError{
			StatusCode: http.StatusBadRequest,
			Error:      "INVALID_INTERNSHIP_ID",
			Details: map[string]any{
				"id": idStr,
			},
		})
		return
	}

	var internshipDto InternshipRequestDto
	apiErr := corehttp.ParseJSONBody(r, &internshipDto)
	if apiErr != nil {
		corehttp.RespondError(w, *apiErr)
		return
	}

	internship, err := h.handlers.Patch.Handle(r.Context(), app.UpdateInternshipCommand{
		ID:          id,
		Title:       internshipDto.Title,
		Label:       internshipDto.Label,
		Description: internshipDto.Description,
		Skills:      internshipDto.Skills,
		Goals:       internshipDto.Goals,
	})

	if err != nil {
		log.Printf("failed to update internship: %v", err)
		corehttp.RespondError(w, mapError(err))
		return
	}

	dto := MapInternship(internship)
	corehttp.Respond(w, http.StatusOK, dto)
}
