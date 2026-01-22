package app

import (
	"context"
	"culand-internship/internal/internship/domain"
)

type GetAllInternshipsHandler struct {
	repo Repository
}

func NewGetAllInternshipsHandler(repo Repository) *GetAllInternshipsHandler {
	return &GetAllInternshipsHandler{repo: repo}
}

func (h *GetAllInternshipsHandler) Handle(ctx context.Context, _ GetAllInternshipsQuery) ([]*domain.Internship, error) {
	return h.repo.GetAll(ctx)
}
