package dto

type CleanBlacklistResponse struct {
	StatusCode int    `json:"status_code"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}

func (_ *CleanBlacklistResponse) CreateCleanBlacklistsResponse(status int, code, message string) *CleanBlacklistResponse {
	return &CleanBlacklistResponse{
		StatusCode: status,
		Code:       code,
		Message:    message,
	}
}
