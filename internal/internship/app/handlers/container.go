package handlers

import (
	"culand-internship/internal/internship/app"
)

type Handlers struct {
	GetAll  *GetAllInternshipsHandler
	GetByID *GetInternshipByIDHandler
	Create  *CreateInternshipHandler
	Update  *UpdateInternshipHandler
	Patch   *PatchInternshipHandler
}

func BuildHandlers(repo app.Repository, txRunner app.TxRunner) *Handlers {
	return &Handlers{
		GetAll:  NewGetAllInternshipsHandler(repo),
		GetByID: NewGetInternshipByIDHandler(repo),
		Create:  NewCreateInternshipHandler(txRunner),
		Update:  NewUpdateInternshipHandler(txRunner),
		Patch:   NewPatchInternshipHandler(txRunner),
	}
}
