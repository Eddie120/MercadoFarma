package codes

type ErrCode string

const (
	InvalidInput        ErrCode = "INVALID_INPUT"
	InternalServerError ErrCode = "INTERNAL_SERVER_ERROR"
)
