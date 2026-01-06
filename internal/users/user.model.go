package users

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	FarmID    uint   `json:"farm_id"` // will be used later for multi-tenancy
	Name      string `json:"name"`
	Email     string `gorm:"unique" json:"email"`
	Password  string `json:"password" binding:"required,min=6"`
	Role      string `json:"role"` // "owner" or "worker"
	CreatedAt time.Time
}
