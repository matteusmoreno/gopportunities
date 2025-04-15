package schemas

import (
	"gorm.io/gorm"
	"time"
)

type Opening struct {
	gorm.Model
	Role     string  `json:"role"`
	Company  string  `json:"company"`
	Location string  `json:"location"`
	Remote   bool    `json:"remote"`
	Link     string  `json:"link"`
	Salary   float64 `json:"salary"`
}

type OpeningResponse struct {
	ID        uint       `json:"id"`
	Role      string     `json:"role"`
	Company   string     `json:"company"`
	Location  string     `json:"location"`
	Remote    bool       `json:"remote"`
	Link      string     `json:"link"`
	Salary    float64    `json:"salary"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func ToOpeningResponse(opening Opening) OpeningResponse {
	var deletedAt *time.Time
	if opening.DeletedAt.Valid {
		deletedAt = &opening.DeletedAt.Time
	}

	return OpeningResponse{
		ID:        opening.ID,
		Role:      opening.Role,
		Company:   opening.Company,
		Location:  opening.Location,
		Remote:    opening.Remote,
		Link:      opening.Link,
		Salary:    opening.Salary,
		CreatedAt: opening.CreatedAt,
		UpdatedAt: opening.UpdatedAt,
		DeletedAt: deletedAt,
	}
}
