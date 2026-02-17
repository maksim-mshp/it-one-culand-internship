package handlers

import (
	"context"
	"culand-internship/internal/internship/app"
	"culand-internship/internal/internship/domain"
)

type CreateInternshipHandler struct {
	txRunner app.TxRunner
}

func NewCreateInternshipHandler(txRunner app.TxRunner) *CreateInternshipHandler {
	return &CreateInternshipHandler{txRunner: txRunner}
}

func (h *CreateInternshipHandler) Handle(ctx context.Context, cmd app.CreateInternshipCommand) (*domain.Internship, error) {
	if cmd.Title == nil {
		return nil, domain.NewMissingFieldError("title")
	}
	if cmd.Label == nil {
		label := ""
		cmd.Label = &label
	}
	if cmd.Description == nil {
		description := ""
		cmd.Description = &description
	}
	if cmd.Skills == nil {
		skills := make([]string, 0)
		cmd.Skills = &skills
	}
	if cmd.Goals == nil {
		goals := make([]string, 0)
		cmd.Goals = &goals
	}

	internship, err := domain.NewInternship(
		*cmd.Title,
		*cmd.Label,
		*cmd.Description,
		*cmd.Skills,
		*cmd.Goals,
	)
	if err != nil {
		return nil, err
	}

	err = h.txRunner.WithinTx(ctx, func(ctx context.Context, repo app.Repository) error {
		id, err := repo.Create(ctx, internship)
		if err != nil {
			return err
		}

		internship, err = repo.GetByID(ctx, id)
		return err
	})
	if err != nil {
		return nil, err
	}

	return internship, nil
}
