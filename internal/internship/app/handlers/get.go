package handlers

import (
	"context"
	"culand-internship/internal/internship/app"
	"culand-internship/internal/internship/domain"
)

type GetAllInternshipsHandler struct {
	repo app.Repository
}

func NewGetAllInternshipsHandler(repo app.Repository) *GetAllInternshipsHandler {
	return &GetAllInternshipsHandler{repo: repo}
}

func (h *GetAllInternshipsHandler) Handle(ctx context.Context, _ app.GetAllInternshipsQuery) ([]*domain.Internship, error) {
	return h.repo.GetAll(ctx)
}

type GetInternshipByIDHandler struct {
	repo app.Repository
}

func NewGetInternshipByIDHandler(repo app.Repository) *GetInternshipByIDHandler {
	return &GetInternshipByIDHandler{repo: repo}
}

func (h *GetInternshipByIDHandler) Handle(ctx context.Context, _ app.GetInternshipByIDQuery, id int) (*domain.Internship, error) {
	return h.repo.GetByID(ctx, id)
}
