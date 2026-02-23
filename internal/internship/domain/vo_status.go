package domain

import "strings"

type Status struct {
	value int
}

const (
	statusUnknown = iota
	statusDraft
	statusActive
	statusArchive
)

var (
	StatusUnknown = Status{value: statusUnknown}
	StatusDraft   = Status{value: statusDraft}
	StatusActive  = Status{value: statusActive}
	StatusArchive = Status{value: statusArchive}
)

func NewStatus(v string) (Status, error) {
	val := strings.ToUpper(strings.TrimSpace(v))
	switch val {
	case "DRAFT":
		return StatusDraft, nil
	case "ACTIVE":
		return StatusActive, nil
	case "ARCHIVE":
		return StatusArchive, nil
	}
	return Status{}, NewInvalidStatusError(val, []string{"DRAFT", "ACTIVE", "ARCHIVE"})
}

func (s Status) Value() string {
	switch s.value {
	case StatusDraft.value:
		return "DRAFT"
	case StatusActive.value:
		return "ACTIVE"
	case StatusArchive.value:
		return "ARCHIVE"
	}
	return "UNKNOWN"
}

func ReconstituteStatus(v string) Status {
	switch strings.ToUpper(v) {
	case "DRAFT":
		return StatusDraft
	case "ACTIVE":
		return StatusActive
	case "ARCHIVE":
		return StatusArchive
	}
	return StatusUnknown
}
