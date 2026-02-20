package domain

type InternshipInfo struct {
	title       Title
	label       Label
	description Description
	skills      Skills
	goals       Goals
}

func NewInternshipInfo(t Title, l Label, d Description, s Skills, g Goals) (InternshipInfo, error) {
	return InternshipInfo{
		title:       t,
		label:       l,
		description: d,
		skills:      s,
		goals:       g,
	}, nil
}

func (i InternshipInfo) Title() Title {
	return i.title
}
func (i InternshipInfo) Label() Label {
	return i.label
}
func (i InternshipInfo) Description() Description {
	return i.description
}
func (i InternshipInfo) Skills() Skills {
	return i.skills
}
func (i InternshipInfo) Goals() Goals {
	return i.goals
}

func ReconstituteInternshipInfo(t Title, l Label, d Description, s Skills, g Goals) InternshipInfo {
	return InternshipInfo{
		title:       t,
		label:       l,
		description: d,
		skills:      s,
		goals:       g,
	}
}
