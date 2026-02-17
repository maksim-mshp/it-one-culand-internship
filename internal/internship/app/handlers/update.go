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

	internship := &domain.Internship{
		ID:          cmd.ID,
		Title:       *cmd.Title,
		Label:       *cmd.Label,
		Description: *cmd.Description,
		Skills:      *cmd.Skills,
		Goals:       *cmd.Goals,
	}
	err := internship.Validate()
	if err != nil {
		return nil, err
	}

	err = h.txRunner.WithinTx(ctx, func(ctx context.Context, repo app.Repository) error {
		err := repo.Update(ctx, internship)
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
		cur, err := repo.GetByIDForUpdate(ctx, cmd.ID)
		if err != nil {
			return err
		}

		if cmd.Title == nil {
			cmd.Title = &cur.Title
		}
		if cmd.Label == nil {
			cmd.Label = &cur.Label
		}
		if cmd.Description == nil {
			cmd.Description = &cur.Description
		}
		if cmd.Skills == nil {
			cmd.Skills = &cur.Skills
		}
		if cmd.Goals == nil {
			cmd.Goals = &cur.Goals
		}

		internship = &domain.Internship{
			ID:          cmd.ID,
			Title:       *cmd.Title,
			Label:       *cmd.Label,
			Description: *cmd.Description,
			Skills:      *cmd.Skills,
			Goals:       *cmd.Goals,
		}

		err = internship.Validate()
		if err != nil {
			return err
		}

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
