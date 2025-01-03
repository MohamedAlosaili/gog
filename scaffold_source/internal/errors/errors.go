package errors

type (
	ErrorProvider interface {
		NewError(c ErrorCode, m string) *AppError
	}

	AppError struct {
		Code    ErrorCode
		Message string
	}
)

func (e *AppError) Error() string {
	return e.Message
}

func Is(err error, code ErrorCode) bool {
	appErr, ok := err.(*AppError)
	return ok && appErr.Code == code
}
