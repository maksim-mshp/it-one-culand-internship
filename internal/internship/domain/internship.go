package domain

type Internship struct {
	id   int
	info InternshipInfo
}

func NewInternship(info InternshipInfo) (Internship, error) {
	return Internship{
		info: info,
	}, nil
}

func (i *Internship) ID() int {
	return i.id
}
func (i *Internship) Info() InternshipInfo {
	return i.info
}

func (i *Internship) SetID(id int) {
	i.id = id
}
func (i *Internship) UpdateInfo(info InternshipInfo) {
	i.info = info
}

func ReconstituteInternship(id int, info InternshipInfo) Internship {
	return Internship{
		id:   id,
		info: info,
	}
}
