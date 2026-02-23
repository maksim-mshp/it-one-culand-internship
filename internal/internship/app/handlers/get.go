package handlers

import (
	"context"
	"culand-internship/internal/internship/app"
	"culand-internship/internal/internship/domain"
)

type GetAllHandler struct {
	repo app.Repository
}

func NewGetAllHandler(repo app.Repository) *GetAllHandler {
	return &GetAllHandler{repo: repo}
}

func (h *GetAllHandler) Handle(ctx context.Context, _ app.GetAllQuery) ([]*domain.Internship, error) {
	return h.repo.GetAllActive(ctx)
}

type GetByIDHandler struct {
	repo app.Repository
}

func NewGetByIDHandler(repo app.Repository) *GetByIDHandler {
	return &GetByIDHandler{repo: repo}
}

func (h *GetByIDHandler) Handle(ctx context.Context, q app.GetByIDQuery) (*domain.Internship, error) {
	internship, err := h.repo.GetByID(ctx, q.ID)
	if err != nil {
		return nil, err
	}
	if internship.Status() != domain.StatusActive {
		return nil, domain.NewNotFoundError(q.ID)
	}
	return internship, nil
}

type GetAllAdminHandler struct {
	repo app.Repository
}

func NewGetAllAdminHandler(repo app.Repository) *GetAllAdminHandler {
	return &GetAllAdminHandler{repo: repo}
}

func (h *GetAllAdminHandler) Handle(ctx context.Context, _ app.GetAllQuery) ([]*domain.Internship, error) {
	return h.repo.GetAll(ctx)
}

type GetByIDAdminHandler struct {
	repo app.Repository
}

func NewGetByIDAdminHandler(repo app.Repository) *GetByIDAdminHandler {
	return &GetByIDAdminHandler{repo: repo}
}

func (h *GetByIDAdminHandler) Handle(ctx context.Context, q app.GetByIDQuery) (*domain.Internship, error) {
	return h.repo.GetByID(ctx, q.ID)
}
