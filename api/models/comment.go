package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	ProjectId uuid.UUID `gorm:"not null;"json:"projectId"`
	Project   Project   `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL;"`
}
