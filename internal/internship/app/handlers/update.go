package handlers

import (
	"context"
	"culand-internship/internal/core/utils"
	"culand-internship/internal/internship/app"
	"culand-internship/internal/internship/domain"
)

type UpdateHandler struct {
	txRunner app.TxRunner
}

func NewUpdateHandler(txRunner app.TxRunner) *UpdateHandler {
	return &UpdateHandler{txRunner: txRunner}
}

func (h *UpdateHandler) Handle(ctx context.Context, cmd app.UpdateCommand) (*domain.Internship, error) {
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

	internship.SetID(cmd.ID)

	err = h.txRunner.WithinTx(ctx, func(ctx context.Context, repo app.Repository) error {
		if err = repo.Update(ctx, internship); err != nil {
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

type PatchHandler struct {
	txRunner app.TxRunner
}

func NewPatchHandler(txRunner app.TxRunner) *PatchHandler {
	return &PatchHandler{txRunner: txRunner}
}

func (h *PatchHandler) Handle(ctx context.Context, cmd app.UpdateCommand) (*domain.Internship, error) {
	internship := &domain.Internship{}

	err := h.txRunner.WithinTx(ctx, func(ctx context.Context, repo app.Repository) error {
		var err error
		internship, err = repo.GetByIDForUpdate(ctx, cmd.ID)
		if err != nil {
			return err
		}

		curInfo := internship.Info()

		title, err := utils.UpdateVO(cmd.Title, domain.NewTitle, curInfo.Title())
		if err != nil {
			return err
		}
		label, err := utils.UpdateVO(cmd.Label, domain.NewLabel, curInfo.Label())
		if err != nil {
			return err
		}
		desc, err := utils.UpdateVO(cmd.Description, domain.NewDescription, curInfo.Description())
		if err != nil {
			return err
		}
		skills, err := utils.UpdateVO(cmd.Skills, domain.NewSkills, curInfo.Skills())
		if err != nil {
			return err
		}
		goals, err := utils.UpdateVO(cmd.Goals, domain.NewGoals, curInfo.Goals())
		if err != nil {
			return err
		}

		info, err := domain.NewInternshipInfo(title, label, desc, skills, goals)
		if err != nil {
			return err
		}

		if err = internship.UpdateInfo(info); err != nil {
			return err
		}

		if err = repo.Update(ctx, internship); err != nil {
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

type UpdateStatusHandler struct {
	txRunner app.TxRunner
}

func NewUpdateStatusHandler(txRunner app.TxRunner) *UpdateStatusHandler {
	return &UpdateStatusHandler{txRunner: txRunner}
}

func (h *UpdateStatusHandler) Handle(ctx context.Context, cmd app.UpdateStatusCommand) (*domain.Internship, error) {
	status, err := utils.BuildVO(cmd.Status, domain.NewStatus)
	if err != nil {
		return nil, err
	}

	internship := &domain.Internship{}

	err = h.txRunner.WithinTx(ctx, func(ctx context.Context, repo app.Repository) error {
		internship, err = repo.GetByIDForUpdate(ctx, cmd.ID)
		if err != nil {
			return err
		}

		internship.SetID(cmd.ID)
		if err = internship.ChangeStatus(status); err != nil {
			return err
		}

		if err = repo.UpdateStatus(ctx, cmd.ID, status); err != nil {
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
