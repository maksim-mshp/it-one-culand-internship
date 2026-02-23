package handlers

import (
	"context"
	"culand-internship/internal/internship/app"
)

type DeleteHandler struct {
	repo app.Repository
}

func NewDeleteHandler(repo app.Repository) *DeleteHandler {
	return &DeleteHandler{repo: repo}
}

func (h *DeleteHandler) Handle(ctx context.Context, cmd app.DeleteCommand) error {
	return h.repo.Delete(ctx, cmd.ID)
}
