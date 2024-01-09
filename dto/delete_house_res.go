package dto

type DeleteHouseResponse struct {
	StatusCode int    `json:"status_code"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}

func (_ *DeleteHouseResponse) CreateDeleteHouseResponse(status int, code, message string) *DeleteHouseResponse {
	return &DeleteHouseResponse{
		StatusCode: status,
		Code:       code,
		Message:    message,
	}
}
