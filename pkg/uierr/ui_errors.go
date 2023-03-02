package uierr

type Error struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
	Details []any  `json:"details"`
}

type ValidationField struct {
	FieldName string `json:"fieldName"`
	Reason    string `json:"reason"`
}

func (e *Error) Error() string {
	return e.Message
}

type Code uint

const (
	CodeSuccess Code = iota
	CodeBadRequest
	CodeResourceNotFound
	CodeInternalServerError
	CodeUnAuthorization
	CodeTokenExpired
	CodeInvalidRequest

	CodeAlreadyExists

	CodeInternal = 999
)

func New(code Code, message string) *Error {
	return &Error{code, message, nil}
}

func Alert(message string) error {
	return &Error{
		CodeBadRequest,
		message,
		nil,
	}
}

func NotFound(message string) error {
	return &Error{
		Code:    CodeResourceNotFound,
		Message: message,
	}
}
func UnAuthorization(message string) error {
	return &Error{
		Code:    CodeUnAuthorization,
		Message: message,
	}
}

func BadInput(field, reason string) error {
	return &Error{
		Code:    CodeInvalidRequest,
		Message: reason,
		Details: []any{
			ValidationField{
				field,
				reason,
			},
		},
	}
}

func Invalid(field, reason string) error {
	return BadInput(field, reason)
}

func AlreadyExist(reason string) error {
	return &Error{
		Code:    CodeAlreadyExists,
		Message: reason,
	}
}

func InternalServer() error {
	return &Error{
		Code:    CodeInternal,
		Message: "internal server error",
	}
}
