package handlers

import (
	"context"
	"culand-internship/internal/internship/app"
	"culand-internship/internal/internship/domain"
)

type UpdateInternshipHandler struct {
	txRunner app.TxRunner
}

func NewUpdateInternshipHandler(txRunner app.TxRunner) *UpdateInternshipHandler {
	return &UpdateInternshipHandler{txRunner: txRunner}
}

func (h *UpdateInternshipHandler) Handle(ctx context.Context, cmd app.UpdateInternshipCommand) (*domain.Internship, error) {
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

	internship.SetID(cmd.ID)

	err = h.txRunner.WithinTx(ctx, func(ctx context.Context, repo app.Repository) error {
		err = repo.Update(ctx, internship)
		if err != nil {
			return err
		}

		internship, err = repo.GetByID(ctx, cmd.ID)
		return err
	})
	if err != nil {
		return nil, err
	}

	return internship, nil
}

type PatchInternshipHandler struct {
	txRunner app.TxRunner
}

func NewPatchInternshipHandler(txRunner app.TxRunner) *PatchInternshipHandler {
	return &PatchInternshipHandler{txRunner: txRunner}
}

func (h *PatchInternshipHandler) Handle(ctx context.Context, cmd app.UpdateInternshipCommand) (*domain.Internship, error) {
	internship := &domain.Internship{}

	err := h.txRunner.WithinTx(ctx, func(ctx context.Context, repo app.Repository) error {
		var err error
		internship, err = repo.GetByIDForUpdate(ctx, cmd.ID)
		if err != nil {
			return err
		}

		title := internship.Info().Title()
		label := internship.Info().Label()
		desc := internship.Info().Description()
		skills := internship.Info().Skills()
		goals := internship.Info().Goals()

		if cmd.Title != nil {
			if title, err = domain.NewTitle(cmd.Title); err != nil {
				return err
			}
		}
		if cmd.Label != nil {
			if label, err = domain.NewLabel(cmd.Label); err != nil {
				return err
			}
		}
		if cmd.Description != nil {
			if desc, err = domain.NewDescription(cmd.Description); err != nil {
				return err
			}
		}
		if cmd.Skills != nil {
			if skills, err = domain.NewSkills(cmd.Skills); err != nil {
				return err
			}
		}
		if cmd.Goals != nil {
			if goals, err = domain.NewGoals(cmd.Goals); err != nil {
				return err
			}
		}

		info, err := domain.NewInternshipInfo(title, label, desc, skills, goals)
		if err != nil {
			return err
		}
		internship.UpdateInfo(info)

		err = repo.Update(ctx, internship)
		if err != nil {
			return err
		}

		internship, err = repo.GetByID(ctx, cmd.ID)
		return err
	})

	if err != nil {
		return nil, err
	}

	return internship, nil
}
