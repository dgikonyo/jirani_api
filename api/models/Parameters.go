package models

import (
	"github.com/google/uuid"
) 

type Parameters struct {
	id uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
}