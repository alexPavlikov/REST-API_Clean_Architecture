package apperror

import "encoding/json"

var (
	ErrorNotFound = NewAppError(nil, "not found", "", "ME-0001")
	NoAuthError   = NewAppError(nil, "unauthorized", "", "ME-0002")
)

type AppError struct {
	Err        error  `json:"-"`
	Message    string `json:"message"`
	DevMessage string `json:"dev_message"`
	Code       string `json:"code"`
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) Marhsal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return marshal
}

func NewAppError(err error, message, devMessage, code string) *AppError {
	return &AppError{
		Err:        err,
		Message:    message,
		DevMessage: devMessage,
		Code:       code,
	}
}

func systemError(err error) *AppError {
	return NewAppError(err, "internal system error", err.Error(), "US-00000")
}
