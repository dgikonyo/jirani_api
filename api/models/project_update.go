package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
) 

type ProjectUpdate struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
}