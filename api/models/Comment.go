package models

import (
	"github.com/google/uuid"
)

type Comment struct {
	id uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	projectId uuid.UUID `gorm:"not null;"json:"projectId"`
	project Project `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL;"`
}