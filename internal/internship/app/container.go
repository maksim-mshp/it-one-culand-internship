package app

type Handlers struct {
	GetAll *GetAllInternshipsHandler
}

func BuildHandlers(repo Repository) *Handlers {
	return &Handlers{
		GetAll: NewGetAllInternshipsHandler(repo),
	}
}
