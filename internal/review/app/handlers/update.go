package handlers

import (
	"context"
	"culand-internship/internal/review/app"
	"culand-internship/internal/review/domain"
	"culand-internship/internal/shared/utils"
)

type UpdateHandler struct {
	txRunner app.TxRunner
}

func NewUpdateHandler(txRunner app.TxRunner) *UpdateHandler {
	return &UpdateHandler{txRunner: txRunner}
}

func (h *UpdateHandler) Handle(ctx context.Context, cmd app.UpdateCommand) (*domain.Review, error) {
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

	review.SetID(cmd.ID)

	err = h.txRunner.WithinTx(ctx, func(ctx context.Context, repo app.Repository) error {
		if err = repo.Update(ctx, review); err != nil {
			return err
		}

		review, err = repo.GetByID(ctx, cmd.ID)
		return err
	})
	if err != nil {
		return nil, err
	}

	return review, nil
}

type PatchHandler struct {
	txRunner app.TxRunner
}

func NewPatchHandler(txRunner app.TxRunner) *PatchHandler {
	return &PatchHandler{txRunner: txRunner}
}

func (h *PatchHandler) Handle(ctx context.Context, cmd app.UpdateCommand) (*domain.Review, error) {
	review := &domain.Review{}

	err := h.txRunner.WithinTx(ctx, func(ctx context.Context, repo app.Repository) error {
		var err error
		review, err = repo.GetByIDForUpdate(ctx, cmd.ID)
		if err != nil {
			return err
		}

		text, err := utils.UpdateVO(cmd.Text, domain.NewText, review.Text())
		if err != nil {
			return err
		}
		author, err := utils.UpdateVO(cmd.Author, domain.NewAuthor, review.Author())
		if err != nil {
			return err
		}
		position, err := utils.UpdateVO(cmd.Position, domain.NewPosition, review.Position())
		if err != nil {
			return err
		}

		*review, err = domain.NewReview(text, author, position)
		if err != nil {
			return err
		}
		review.SetID(cmd.ID)

		if err = repo.Update(ctx, review); err != nil {
			return err
		}

		review, err = repo.GetByID(ctx, cmd.ID)
		return err
	})

	if err != nil {
		return nil, err
	}

	return review, nil
}
