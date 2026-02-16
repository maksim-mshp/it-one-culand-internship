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

type GetInternshipByIDHandler struct {
	repo Repository
}

func NewGetInternshipByIDHandler(repo Repository) *GetInternshipByIDHandler {
	return &GetInternshipByIDHandler{repo: repo}
}

func (h *GetInternshipByIDHandler) Handle(ctx context.Context, _ GetInternshipByIDQuery, id int) (*domain.Internship, error) {
	return h.repo.GetByID(ctx, id)
}

type CreateInternshipHandler struct {
	repo Repository
}

func NewCreateInternshipHandler(repo Repository) *CreateInternshipHandler {
	return &CreateInternshipHandler{repo: repo}
}

func (h *CreateInternshipHandler) Handle(ctx context.Context, cmd CreateInternshipCommand) (*domain.Internship, error) {
	internship, err := domain.NewInternship(
		cmd.Title,
		cmd.Label,
		cmd.Description,
		cmd.Skills,
		cmd.Goals,
	)
	if err != nil {
		return nil, err
	}

	id, err := h.repo.Create(ctx, internship)
	if err != nil {
		return nil, err
	}

	return h.repo.GetByID(ctx, id)
}
