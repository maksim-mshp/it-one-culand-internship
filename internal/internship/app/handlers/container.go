package handlers

import (
	"culand-internship/internal/internship/app"
)

type Handlers struct {
	GetAll       *GetAllHandler
	GetByID      *GetByIDHandler
	Create       *CreateHandler
	Update       *UpdateHandler
	Patch        *PatchHandler
	GetAllAdmin  *GetAllAdminHandler
	GetByIDAdmin *GetByIDAdminHandler
	UpdateStatus *UpdateStatusHandler
}

func BuildHandlers(repo app.Repository, txRunner app.TxRunner) *Handlers {
	return &Handlers{
		GetAll:       NewGetAllHandler(repo),
		GetByID:      NewGetByIDHandler(repo),
		Create:       NewCreateHandler(txRunner),
		Update:       NewUpdateHandler(txRunner),
		Patch:        NewPatchHandler(txRunner),
		GetAllAdmin:  NewGetAllAdminHandler(repo),
		GetByIDAdmin: NewGetByIDAdminHandler(repo),
		UpdateStatus: NewUpdateStatusHandler(txRunner),
	}
}
