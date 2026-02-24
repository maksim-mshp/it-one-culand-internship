package handlers

import "culand-internship/internal/review/app"

type Handlers struct {
	GetAll  *GetAllHandler
	GetByID *GetByIDHandler
	Create  *CreateHandler
	Update  *UpdateHandler
	Patch   *PatchHandler
	Delete  *DeleteHandler
}

func BuildHandlers(repo app.Repository, txRunner app.TxRunner) *Handlers {
	return &Handlers{
		GetAll:  NewGetAllHandler(repo),
		GetByID: NewGetByIDHandler(repo),
		Create:  NewCreateHandler(txRunner),
		Update:  NewUpdateHandler(txRunner),
		Patch:   NewPatchHandler(txRunner),
		Delete:  NewDeleteHandler(repo),
	}
}
