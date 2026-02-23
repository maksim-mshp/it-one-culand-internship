package handlers

import (
	"context"
	"culand-internship/internal/faq/app"
	"culand-internship/internal/faq/domain"
	"culand-internship/internal/shared/utils"
)

type UpdateHandler struct {
	txRunner app.TxRunner
}

func NewUpdateHandler(txRunner app.TxRunner) *UpdateHandler {
	return &UpdateHandler{txRunner: txRunner}
}

func (h *UpdateHandler) Handle(ctx context.Context, cmd app.UpdateCommand) (*domain.FAQ, error) {
	question, err := utils.BuildVO(cmd.Question, domain.NewQuestion)
	if err != nil {
		return nil, err
	}
	answer, err := utils.BuildVO(cmd.Answer, domain.NewAnswer)
	if err != nil {
		return nil, err
	}

	val, err := domain.NewFAQ(question, answer)
	if err != nil {
		return nil, err
	}
	faq := &val

	faq.SetID(cmd.ID)

	err = h.txRunner.WithinTx(ctx, func(ctx context.Context, repo app.Repository) error {
		if err = repo.Update(ctx, faq); err != nil {
			return err
		}

		faq, err = repo.GetByID(ctx, cmd.ID)
		return err
	})
	if err != nil {
		return nil, err
	}

	return faq, nil
}

type PatchHandler struct {
	txRunner app.TxRunner
}

func NewPatchHandler(txRunner app.TxRunner) *PatchHandler {
	return &PatchHandler{txRunner: txRunner}
}

func (h *PatchHandler) Handle(ctx context.Context, cmd app.UpdateCommand) (*domain.FAQ, error) {
	faq := &domain.FAQ{}

	err := h.txRunner.WithinTx(ctx, func(ctx context.Context, repo app.Repository) error {
		var err error
		faq, err = repo.GetByIDForUpdate(ctx, cmd.ID)
		if err != nil {
			return err
		}

		question, err := utils.UpdateVO(cmd.Question, domain.NewQuestion, faq.Question())
		if err != nil {
			return err
		}
		answer, err := utils.UpdateVO(cmd.Answer, domain.NewAnswer, faq.Answer())
		if err != nil {
			return err
		}

		*faq, err = domain.NewFAQ(question, answer)
		if err != nil {
			return err
		}
		faq.SetID(cmd.ID)

		if err = repo.Update(ctx, faq); err != nil {
			return err
		}

		faq, err = repo.GetByID(ctx, cmd.ID)
		return err
	})

	if err != nil {
		return nil, err
	}

	return faq, nil
}
