package models

type UrgentCase struct {
	Id uint64 `json:"id" `
	Title string `json:"title"`
	Description string `json:"description"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
