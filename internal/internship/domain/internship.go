package domain

type Internship struct {
	id     int
	info   InternshipInfo
	status Status
}

func NewInternship(info InternshipInfo) (Internship, error) {
	return Internship{
		info:   info,
		status: StatusDraft,
	}, nil
}

func (i *Internship) ID() int {
	return i.id
}
func (i *Internship) Info() InternshipInfo {
	return i.info
}
func (i *Internship) Status() Status {
	return i.status
}

func (i *Internship) SetID(id int) {
	i.id = id
}
func (i *Internship) UpdateInfo(info InternshipInfo) error {
	i.info = info
	return nil
}
func (i *Internship) ChangeStatus(status Status) error {
	i.status = status
	return nil
}

func ReconstituteInternship(id int, info InternshipInfo, status Status) Internship {
	return Internship{
		id:     id,
		info:   info,
		status: status,
	}
}
