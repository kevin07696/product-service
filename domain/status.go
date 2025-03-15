package domain

type StatusCode uint8

const (
	StatusOK              StatusCode = 0
	StatusInvalidArgument StatusCode = 1
	StatusUnauthorized    StatusCode = 2
	StatusNotFound        StatusCode = 3
	StatusAlreadyExists   StatusCode = 4
	StatusInternalError   StatusCode = 5
)
