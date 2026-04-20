package domain

type Status int

const (
	Unknown Status = iota
	Pending
	Active
	Inactive
	Banned
)

//nolint:exhaustive
func (s Status) String() string {
	switch s {
	case Pending:
		return "pending"
	case Active:
		return "active"
	case Inactive:
		return "inactive"
	case Banned:
		return "banned"
	default:
		return "unknown"
	}
}
