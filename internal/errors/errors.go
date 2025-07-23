package errors

type AppError struct {
	Code    string
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func New(code, msg string) *AppError {
	return &AppError{
		Code:    code,
		Message: msg,
	}
}

func Wrap(code, msg string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: msg,
		Err:     err,
	}
}

var (
	ErrNotFound       = New("NOT_FOUND", "Resource not found")
	ErrUnauthorized   = New("UNAUTHORIZED", "Unauthorized access")
	ErrValidation     = New("VALIDATION_ERROR", "Validation error")
	ErrInternalServer = New("INTERNAL_SERVER_ERROR", "Internal server error")
)
