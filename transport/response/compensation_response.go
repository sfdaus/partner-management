package response

// Get List Response
type GetListCompensationRes struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

type CreateCompensationRes struct {
	ID string `json:"id"`
}
