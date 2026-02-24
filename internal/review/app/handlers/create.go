package handlers

import (
	"context"
	"culand-internship/internal/review/app"
	"culand-internship/internal/review/domain"
	"culand-internship/internal/shared/utils"
)

type CreateHandler struct {
	txRunner app.TxRunner
}

func NewCreateHandler(txRunner app.TxRunner) *CreateHandler {
	return &CreateHandler{txRunner: txRunner}
}

func (h *CreateHandler) Handle(ctx context.Context, cmd app.CreateCommand) (*domain.Review, error) {
	text, err := utils.BuildVO(cmd.Text, domain.NewText)
	if err != nil {
		return nil, err
	}
	author, err := utils.BuildVO(cmd.Author, domain.NewAuthor)
	if err != nil {
		return nil, err
	}
	position, err := utils.BuildVO(cmd.Position, domain.NewPosition)
	if err != nil {
		return nil, err
	}

	val, err := domain.NewReview(text, author, position)
	if err != nil {
		return nil, err
	}
	review := &val

	err = h.txRunner.WithinTx(ctx, func(ctx context.Context, repo app.Repository) error {
		id, err := repo.Create(ctx, review)
		if err != nil {
			return err
		}

		review, err = repo.GetByID(ctx, id)
		return err
	})
	if err != nil {
		return nil, err
	}

	return review, nil
}
