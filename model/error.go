package model

type ErrNotFound struct {
	Resource string
	ID       int
}

func (e ErrNotFound) Error() string {
	return "not found"
}
