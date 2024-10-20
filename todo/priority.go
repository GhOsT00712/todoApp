package todo

type Priority int

const (
	LOW Priority = iota
	MEDIUM
	HIGH
)

func (s Priority) String() string {
	switch s {
	case LOW:
		return "LOW"
	case MEDIUM:
		return "MEDIUM"
	case HIGH:
		return "HIGH"
	default:
		return "Invalid Priority"
	}
}
