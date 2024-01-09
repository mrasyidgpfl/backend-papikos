package dto

type UpdatePickUpStatusRequest struct {
	PickUpID uint   `json:"pick_up_id"`
	Status   string `json:"status"`
}
