package handlers

import (
	"context"
	"culand-internship/internal/faq/app"
	"culand-internship/internal/faq/domain"
)

type GetAllHandler struct {
	repo app.Repository
}

func NewGetAllHandler(repo app.Repository) *GetAllHandler {
	return &GetAllHandler{repo: repo}
}

func (h *GetAllHandler) Handle(ctx context.Context, _ app.GetAllQuery) ([]*domain.FAQ, error) {
	return h.repo.GetAll(ctx)
}

type GetByIDHandler struct {
	repo app.Repository
}

func NewGetByIDHandler(repo app.Repository) *GetByIDHandler {
	return &GetByIDHandler{repo: repo}
}

func (h *GetByIDHandler) Handle(ctx context.Context, q app.GetByIDQuery) (*domain.FAQ, error) {
	return h.repo.GetByID(ctx, q.ID)
}
