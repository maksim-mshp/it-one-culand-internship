package app

type Handlers struct {
	GetAll  *GetAllInternshipsHandler
	GetByID *GetInternshipByIDHandler
}

func BuildHandlers(repo Repository) *Handlers {
	return &Handlers{
		GetAll:  NewGetAllInternshipsHandler(repo),
		GetByID: NewGetInternshipByIDHandler(repo),
	}
}
