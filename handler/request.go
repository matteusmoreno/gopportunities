package handler

// Create Opening Request
type CreateOpeningRequest struct {
	Role     string  `json:"role" binding:"required"`
	Company  string  `json:"company" binding:"required"`
	Location string  `json:"location" binding:"required"`
	Remote   *bool   `json:"remote" binding:"required"`
	Link     string  `json:"link" binding:"required"`
	Salary   float64 `json:"salary" binding:"required"`
}

type UpdateOpeningRequest struct {
	ID       uint    `json:"id" binding:"required"`
	Role     string  `json:"role"`
	Company  string  `json:"company"`
	Location string  `json:"location"`
	Remote   *bool   `json:"remote"`
	Link     string  `json:"link"`
	Salary   float64 `json:"salary"`
}
