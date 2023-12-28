package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	ProjectID uuid.UUID `gorm:"not null;" json:"project_id"`
}
