package handlers

import (
	"context"
	"culand-internship/internal/review/app"
	"culand-internship/internal/review/domain"
)

type GetAllHandler struct {
	repo app.Repository
}

func NewGetAllHandler(repo app.Repository) *GetAllHandler {
	return &GetAllHandler{repo: repo}
}

func (h *GetAllHandler) Handle(ctx context.Context, _ app.GetAllQuery) ([]*domain.Review, error) {
	return h.repo.GetAll(ctx)
}

type GetByIDHandler struct {
	repo app.Repository
}

func NewGetByIDHandler(repo app.Repository) *GetByIDHandler {
	return &GetByIDHandler{repo: repo}
}

func (h *GetByIDHandler) Handle(ctx context.Context, q app.GetByIDQuery) (*domain.Review, error) {
	return h.repo.GetByID(ctx, q.ID)
}
