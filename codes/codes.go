package codes

type ErrCode string

const (
	InvalidInput        ErrCode = "INVALID_INPUT"
	InternalServerError ErrCode = "INTERNAL_SERVER_ERROR"
	Unauthorized        ErrCode = "INVALID_AUTHENTICATION"
	DataBaseError       ErrCode = "DATABASE_ERROR"
	ResourceNotFound    ErrCode = "RESOURCE_NOT_FOUND"
)
