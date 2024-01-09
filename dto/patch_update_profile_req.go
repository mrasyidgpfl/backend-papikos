package dto

type UpdateProfileRequest struct {
	ID       uint   `json:"user_id"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
	Address  string `json:"address"`
	CityID   uint   `json:"city_id"`
	Role     string `json:"role"`
}
