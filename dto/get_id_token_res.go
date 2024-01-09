package dto

type SignInResponse struct {
	StatusCode int     `json:"status_code"`
	Code       string  `json:"code"`
	Message    string  `json:"message"`
	JWTToken   IDToken `json:"data"`
}

type IDToken struct {
	IDToken string `json:"id_token"`
}

func (tr *SignInResponse) CreateSignInResponse(status int, code, message, token string) *SignInResponse {
	return &SignInResponse{
		StatusCode: status,
		Code:       code,
		Message:    message,
		JWTToken:   IDToken{token},
	}
}
