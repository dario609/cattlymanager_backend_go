package animals

import "time"

type Animal struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	EarTag       string    `json:"ear_tag"`
	ElectronicID string    `json:"electronic_id"`
	Name         string    `json:"name"`
	Breed        string    `json:"breed"`
	Sex          string    `json:"sex"` // Bull, Cow, Steer, Heifer
	BirthDate    time.Time `json:"birth_date"`
	BirthWeight  float64   `json:"birth_weight"`

	LocationID  uint   `json:"location_id"` // For movement module
	GroupID     uint   `json:"group_id"`    // grouping / herd id
	EarTagColor string `json:"ear_tag_color"`
	Status      string `json:"status"` // Active, Sold, Dead

	CreatedAt time.Time
	UpdatedAt time.Time
}
