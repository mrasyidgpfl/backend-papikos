package dto

type SignUpRequest struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
	Address  string `json:"address"`
	CityID   uint   `json:"city_id"`
}
