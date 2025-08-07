package response

import "time"

// Create Response
type CreatePartnerRes struct {
	ID string `json:"id"`
}

// Get List Response
type GetListPartnerRes struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

// Get Detail Response
type GetDetailPartnerRes struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   string    `json:"updated_by"`
}
