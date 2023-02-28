package uierr

type Error struct {
	code    Code
	message string
	details interface{}
}

type ValidationField struct {
	FieldName string `json:"fieldName"`
	Reason    string `json:"reason"`
}

func (e Error) Error() string {
	return e.message
}
func (e Error) Code() Code {
	return e.code
}

func (e Error) Message() string {
	return e.message
}

type Code uint

const (
	CodeSuccess = iota
	CodeBadRequest
	CodeResourceNotFound
	CodeInternalServerError
	CodeUnAuthorization
	CodeTokenExpired
	CodeInvalidRequest

	CodeAlreadyExists
)

func New(code Code, message string) Error {
	return Error{code: code, message: message}
}

func Alert(message string) Error {
	return Error{
		code:    CodeBadRequest,
		message: message}
}

func NotFound(message string) Error {
	return Error{code: CodeResourceNotFound, message: message}
}
func UnAuthorization(message string) Error {
	return Error{
		code:    CodeUnAuthorization,
		message: message,
	}
}

func BadInput(field, reason string) Error {
	return Error{
		code:    CodeInvalidRequest,
		message: reason,
		details: []ValidationField{
			{
				field,
				reason,
			},
		},
	}
}

func AlreadyExist(reason string) error {
	return Error{
		code:    CodeAlreadyExists,
		message: reason,
	}
}
