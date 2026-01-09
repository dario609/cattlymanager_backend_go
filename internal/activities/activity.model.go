package activities

import (
	"livestock-saas/server/internal/animals"
	"time"
)

// Activity represents an event associated with an animal, such as a Sale or Death record.
type Activity struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	AnimalID uint      `json:"animal_id" gorm:"index"`
	Animal   *animals.Animal `json:"animal,omitempty" gorm:"foreignKey:AnimalID"`
	Type     string    `json:"type"` // activity type: "sale" or "dead"
	Date     time.Time `json:"date"`

	// Sale-specific fields
	SalePrice    float64 `json:"sale_price,omitempty"`
	CustomerName string  `json:"customer_name,omitempty"`

	// Death-specific fields
	DeathCause string `json:"death_cause,omitempty"`

	Notes string `json:"notes,omitempty" gorm:"type:text"`

	// AttachmentURLs stores attachments as a JSON string or comma-separated list.
	// Keep simple here; serializers or a separate attachments table can be added later.
	AttachmentURLs string `json:"attachment_urls,omitempty" gorm:"type:text"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

const (
	ActivityTypeSale = "sale"
	ActivityTypeDead = "dead"
)
