package app

type Handlers struct {
	GetAll  *GetAllInternshipsHandler
	GetByID *GetInternshipByIDHandler
	Create  *CreateInternshipHandler
}

func BuildHandlers(repo Repository) *Handlers {
	return &Handlers{
		GetAll:  NewGetAllInternshipsHandler(repo),
		GetByID: NewGetInternshipByIDHandler(repo),
		Create:  NewCreateInternshipHandler(repo),
	}
}
