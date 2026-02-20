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
	title, err := domain.NewTitle(cmd.Title)
	if err != nil {
		return nil, err
	}
	label, err := domain.NewLabel(cmd.Label)
	if err != nil {
		return nil, err
	}
	desc, err := domain.NewDescription(cmd.Description)
	if err != nil {
		return nil, err
	}
	skills, err := domain.NewSkills(cmd.Skills)
	if err != nil {
		return nil, err
	}
	goals, err := domain.NewGoals(cmd.Goals)
	if err != nil {
		return nil, err
	}

	info, err := domain.NewInternshipInfo(title, label, desc, skills, goals)
	if err != nil {
		return nil, err
	}

	i, err := domain.NewInternship(info)
	if err != nil {
		return nil, err
	}
	internship := &i

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
