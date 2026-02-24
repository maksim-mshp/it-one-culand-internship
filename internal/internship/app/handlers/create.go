package handlers

import (
	"context"
	"culand-internship/internal/internship/app"
	"culand-internship/internal/internship/domain"
	"culand-internship/internal/shared/utils"
)

type CreateHandler struct {
	txRunner app.TxRunner
}

func NewCreateHandler(txRunner app.TxRunner) *CreateHandler {
	return &CreateHandler{txRunner: txRunner}
}

func (h *CreateHandler) Handle(ctx context.Context, cmd app.CreateCommand) (*domain.Internship, error) {
	title, err := utils.BuildVO(cmd.Title, domain.NewTitle)
	if err != nil {
		return nil, err
	}
	label, err := utils.BuildVO(cmd.Label, domain.NewLabel)
	if err != nil {
		return nil, err
	}
	desc, err := utils.BuildVO(cmd.Description, domain.NewDescription)
	if err != nil {
		return nil, err
	}
	skills, err := utils.BuildVO(cmd.Skills, domain.NewSkills)
	if err != nil {
		return nil, err
	}
	goals, err := utils.BuildVO(cmd.Goals, domain.NewGoals)
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
