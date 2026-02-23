package v1

import (
	corehttp "culand-internship/internal/core/http"
	"culand-internship/internal/faq/app"
	"culand-internship/internal/faq/app/handlers"
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

// @Summary		Получить все часто задаваемые вопросы
// @Tags		faq
// @Success		200 {array} FAQDto
// @Router		/faq [GET]
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	all, err := h.handlers.GetAll.Handle(r.Context(), app.GetAllQuery{})
	if err != nil {
		log.Printf("failed to get faq: %v", err)
		corehttp.RespondError(w, mapError(err))
		return
	}

	allDto := make([]FAQDto, 0, len(all))
	for _, i := range all {
		allDto = append(allDto, MapFAQ(i))
	}

	corehttp.Respond(w, http.StatusOK, allDto)
}

// @Summary		Получить часто задаваемый вопрос по ID
// @Tags		faq
// @Param		id path int true "ID"
// @Success		200 {object} FAQDto
// @Failure		400 {object} APIError
// @Failure		404 {object} APIError
// @Router		/faq/{id} [GET]
func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		corehttp.RespondError(w, corehttp.APIError{
			StatusCode: http.StatusBadRequest,
			Error:      "INVALID_FAQ_ID",
			Details: map[string]any{
				"id": idStr,
			},
		})
		return
	}

	faq, err := h.handlers.GetByID.Handle(r.Context(), app.GetByIDQuery{ID: id})
	if err != nil {
		log.Printf("failed to get faq: %v", err)
		corehttp.RespondError(w, mapError(err))
		return
	}

	dto := MapFAQ(faq)
	corehttp.Respond(w, http.StatusOK, dto)
}

// @Summary		Создать часто задаваемый вопрос
// @Tags		admin
// @Param		faq body FAQRequestDto true "Часто задаваемый вопрос"
// @Success		201 {object} FAQDto
// @Failure		401 {object} APIError
// @Failure		422 {object} APIError
// @Router		/admin/faq [POST]
// @Security	Bearer
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var faqDto FAQRequestDto
	apiErr := corehttp.ParseJSONBody(r, &faqDto)
	if apiErr != nil {
		corehttp.RespondError(w, *apiErr)
		return
	}

	faq, err := h.handlers.Create.Handle(r.Context(), app.CreateCommand{
		Question: faqDto.Question,
		Answer:   faqDto.Answer,
	})
	if err != nil {
		log.Printf("failed to create faq: %v", err)
		corehttp.RespondError(w, mapError(err))
		return
	}

	dto := MapFAQ(faq)
	location := fmt.Sprintf("/api/v1/faq/%d", faq.ID())
	w.Header().Set("Location", location)
	corehttp.Respond(w, http.StatusCreated, dto)
}

// @Summary		Обновить часто задаваемый вопрос
// @Tags		admin
// @Param		id path int true "ID"
// @Param		faq body FAQRequestDto true "Часто задаваемый вопрос"
// @Success		200 {object} FAQDto
// @Failure		400 {object} APIError
// @Failure		401 {object} APIError
// @Failure		404 {object} APIError
// @Failure		422 {object} APIError
// @Router		/admin/faq/{id} [PUT]
// @Security	Bearer
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		corehttp.RespondError(w, corehttp.APIError{
			StatusCode: http.StatusBadRequest,
			Error:      "INVALID_FAQ_ID",
			Details: map[string]any{
				"id": idStr,
			},
		})
		return
	}

	var faqDto FAQRequestDto
	apiErr := corehttp.ParseJSONBody(r, &faqDto)
	if apiErr != nil {
		corehttp.RespondError(w, *apiErr)
		return
	}

	faq, err := h.handlers.Update.Handle(r.Context(), app.UpdateCommand{
		ID:       id,
		Question: faqDto.Question,
		Answer:   faqDto.Answer,
	})

	if err != nil {
		log.Printf("failed to update faq: %v", err)
		corehttp.RespondError(w, mapError(err))
		return
	}

	dto := MapFAQ(faq)
	corehttp.Respond(w, http.StatusOK, dto)
}

// @Summary		Обновить часто задаваемый вопрос
// @Tags		admin
// @Param		id path int true "ID"
// @Param		faq body FAQRequestDto true "Часто задаваемый вопрос"
// @Success		200 {object} FAQDto
// @Failure		400 {object} APIError
// @Failure		401 {object} APIError
// @Failure		404 {object} APIError
// @Failure		422 {object} APIError
// @Router		/admin/faq/{id} [PATCH]
// @Security	Bearer
func (h *Handler) Patch(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		corehttp.RespondError(w, corehttp.APIError{
			StatusCode: http.StatusBadRequest,
			Error:      "INVALID_FAQ_ID",
			Details: map[string]any{
				"id": idStr,
			},
		})
		return
	}

	var faqDto FAQRequestDto
	apiErr := corehttp.ParseJSONBody(r, &faqDto)
	if apiErr != nil {
		corehttp.RespondError(w, *apiErr)
		return
	}

	faq, err := h.handlers.Patch.Handle(r.Context(), app.UpdateCommand{
		ID:       id,
		Question: faqDto.Question,
		Answer:   faqDto.Answer,
	})

	if err != nil {
		log.Printf("failed to update faq: %v", err)
		corehttp.RespondError(w, mapError(err))
		return
	}

	dto := MapFAQ(faq)
	corehttp.Respond(w, http.StatusOK, dto)
}

// @Summary		Удалить часто задаваемый вопрос
// @Tags		admin
// @Param		id path int true "ID"
// @Success		204
// @Failure		400 {object} APIError
// @Failure		404 {object} APIError
// @Router		/admin/faq/{id} [DELETE]
// @Security	Bearer
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		corehttp.RespondError(w, corehttp.APIError{
			StatusCode: http.StatusBadRequest,
			Error:      "INVALID_FAQ_ID",
			Details: map[string]any{
				"id": idStr,
			},
		})
		return
	}

	err = h.handlers.Delete.Handle(r.Context(), app.DeleteCommand{ID: id})
	if err != nil {
		log.Printf("failed to delete faq: %v", err)
		corehttp.RespondError(w, mapError(err))
		return
	}

	corehttp.Respond(w, http.StatusNoContent, nil)
}
