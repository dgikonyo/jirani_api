package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
) 

type Parameters struct {
	gorm.Model
	id uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
}