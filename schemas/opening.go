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
