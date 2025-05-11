package response

const (
	ErrCodeSuccess      = 20001
	ErrCodeParamInvalid = 20002

	ErrInvalidToken = 20003
)

var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "param invalid",
	ErrInvalidToken:     "invalid token",
}
