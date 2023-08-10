package xerr

type Error struct {
	Message	string	`json:"message"`
	Code	int		`json:"code"`
}

var (
	NoError = &Error{Message: "ok", Code: 200}
	Internal = &Error{Message: "internal server error", Code: 500}
	AuthenticationFailed = &Error{Message: "authentication failed", Code: 1001}
	InvalidParameter = &Error{Message: "invalid parameter", Code: 1002}
)

func (e *Error)String() string {
	return e.Message
}

func (e *Error)Error() string {
	return e.String()
}

func New(err error) *Error {
	// 校正认证错误信息
	if err == nil {err = NoError}

	// 初始化错误对象
	switch err.(type) {
	case *Error: return err.(*Error)
	default: return &Error{Code: Internal.Code, Message: err.Error()}
	}
}
