package models

import (
	"github.com/google/uuid"
) 

type InvestmentOptionCatalog struct {
	id uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
}