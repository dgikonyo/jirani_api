package models

import (
	"github.com/google/uuid"
) 

type Organization struct {
	id uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
}