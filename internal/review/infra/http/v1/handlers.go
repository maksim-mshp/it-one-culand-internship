package v1

import (
	corehttp "culand-internship/internal/core/http"
	"culand-internship/internal/review/app"
	"culand-internship/internal/review/app/handlers"
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

// @Summary		Получить все отзывы
// @Tags		review
// @Success		200 {array} ReviewDto
// @Router		/review [GET]
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	all, err := h.handlers.GetAll.Handle(r.Context(), app.GetAllQuery{})
	if err != nil {
		log.Printf("failed to get review: %v", err)
		corehttp.RespondError(w, mapError(err))
		return
	}

	allDto := make([]ReviewDto, 0, len(all))
	for _, i := range all {
		allDto = append(allDto, MapReview(i))
	}

	corehttp.Respond(w, http.StatusOK, allDto)
}

// @Summary		Получить отзыв по ID
// @Tags		review
// @Param		id path int true "ID"
// @Success		200 {object} ReviewDto
// @Failure		400 {object} APIError
// @Failure		404 {object} APIError
// @Router		/review/{id} [GET]
func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		corehttp.RespondError(w, corehttp.APIError{
			StatusCode: http.StatusBadRequest,
			Error:      "INVALID_REVIEW_ID",
			Details: map[string]any{
				"id": idStr,
			},
		})
		return
	}

	review, err := h.handlers.GetByID.Handle(r.Context(), app.GetByIDQuery{ID: id})
	if err != nil {
		log.Printf("failed to get review: %v", err)
		corehttp.RespondError(w, mapError(err))
		return
	}

	dto := MapReview(review)
	corehttp.Respond(w, http.StatusOK, dto)
}

// @Summary		Создать отзыв
// @Tags		admin
// @Param		review body ReviewRequestDto true "Отзыв"
// @Success		201 {object} ReviewDto
// @Failure		401 {object} APIError
// @Failure		422 {object} APIError
// @Router		/admin/review [POST]
// @Security	Bearer
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var reviewDto ReviewRequestDto
	apiErr := corehttp.ParseJSONBody(r, &reviewDto)
	if apiErr != nil {
		corehttp.RespondError(w, *apiErr)
		return
	}

	review, err := h.handlers.Create.Handle(r.Context(), app.CreateCommand{
		Text:     reviewDto.Text,
		Author:   reviewDto.Author,
		Position: reviewDto.Position,
	})
	if err != nil {
		log.Printf("failed to create review: %v", err)
		corehttp.RespondError(w, mapError(err))
		return
	}

	dto := MapReview(review)
	location := fmt.Sprintf("/api/v1/review/%d", review.ID())
	w.Header().Set("Location", location)
	corehttp.Respond(w, http.StatusCreated, dto)
}

// @Summary		Обновить отзыв
// @Tags		admin
// @Param		id path int true "ID"
// @Param		review body ReviewRequestDto true "Отзыв"
// @Success		200 {object} ReviewDto
// @Failure		400 {object} APIError
// @Failure		401 {object} APIError
// @Failure		404 {object} APIError
// @Failure		422 {object} APIError
// @Router		/admin/review/{id} [PUT]
// @Security	Bearer
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		corehttp.RespondError(w, corehttp.APIError{
			StatusCode: http.StatusBadRequest,
			Error:      "INVALID_REVIEW_ID",
			Details: map[string]any{
				"id": idStr,
			},
		})
		return
	}

	var reviewDto ReviewRequestDto
	apiErr := corehttp.ParseJSONBody(r, &reviewDto)
	if apiErr != nil {
		corehttp.RespondError(w, *apiErr)
		return
	}

	review, err := h.handlers.Update.Handle(r.Context(), app.UpdateCommand{
		ID:       id,
		Text:     reviewDto.Text,
		Author:   reviewDto.Author,
		Position: reviewDto.Position,
	})
	if err != nil {
		log.Printf("failed to update review: %v", err)
		corehttp.RespondError(w, mapError(err))
		return
	}

	dto := MapReview(review)
	corehttp.Respond(w, http.StatusOK, dto)
}

// @Summary		Обновить отзыв
// @Tags		admin
// @Param		id path int true "ID"
// @Param		review body ReviewRequestDto true "Отзыв"
// @Success		200 {object} ReviewDto
// @Failure		400 {object} APIError
// @Failure		401 {object} APIError
// @Failure		404 {object} APIError
// @Failure		422 {object} APIError
// @Router		/admin/review/{id} [PATCH]
// @Security	Bearer
func (h *Handler) Patch(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		corehttp.RespondError(w, corehttp.APIError{
			StatusCode: http.StatusBadRequest,
			Error:      "INVALID_REVIEW_ID",
			Details: map[string]any{
				"id": idStr,
			},
		})
		return
	}

	var reviewDto ReviewRequestDto
	apiErr := corehttp.ParseJSONBody(r, &reviewDto)
	if apiErr != nil {
		corehttp.RespondError(w, *apiErr)
		return
	}

	review, err := h.handlers.Patch.Handle(r.Context(), app.UpdateCommand{
		ID:       id,
		Text:     reviewDto.Text,
		Author:   reviewDto.Author,
		Position: reviewDto.Position,
	})
	if err != nil {
		log.Printf("failed to update review: %v", err)
		corehttp.RespondError(w, mapError(err))
		return
	}

	dto := MapReview(review)
	corehttp.Respond(w, http.StatusOK, dto)
}

// @Summary		Удалить отзыв
// @Tags		admin
// @Param		id path int true "ID"
// @Success		204
// @Failure		400 {object} APIError
// @Failure		404 {object} APIError
// @Router		/admin/review/{id} [DELETE]
// @Security	Bearer
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		corehttp.RespondError(w, corehttp.APIError{
			StatusCode: http.StatusBadRequest,
			Error:      "INVALID_REVIEW_ID",
			Details: map[string]any{
				"id": idStr,
			},
		})
		return
	}

	err = h.handlers.Delete.Handle(r.Context(), app.DeleteCommand{ID: id})
	if err != nil {
		log.Printf("failed to delete review: %v", err)
		corehttp.RespondError(w, mapError(err))
		return
	}

	corehttp.Respond(w, http.StatusNoContent, nil)
}
