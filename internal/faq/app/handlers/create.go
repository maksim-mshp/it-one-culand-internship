package handlers

import (
	"context"
	"culand-internship/internal/faq/app"
	"culand-internship/internal/faq/domain"
	"culand-internship/internal/shared/utils"
)

type CreateHandler struct {
	txRunner app.TxRunner
}

func NewCreateHandler(txRunner app.TxRunner) *CreateHandler {
	return &CreateHandler{txRunner: txRunner}
}

func (h *CreateHandler) Handle(ctx context.Context, cmd app.CreateCommand) (*domain.FAQ, error) {
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

	err = h.txRunner.WithinTx(ctx, func(ctx context.Context, repo app.Repository) error {
		id, err := repo.Create(ctx, faq)
		if err != nil {
			return err
		}

		faq, err = repo.GetByID(ctx, id)
		return err
	})
	if err != nil {
		return nil, err
	}

	return faq, nil
}
