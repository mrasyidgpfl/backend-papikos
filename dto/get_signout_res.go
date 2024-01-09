package dto

type SignOutResponse struct {
	StatusCode int    `json:"status_code"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}

func (_ *SignOutResponse) CreateSignOutResponse(status int, code, message string) *SignOutResponse {
	return &SignOutResponse{
		StatusCode: status,
		Code:       code,
		Message:    message,
	}
}
