package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/google/uuid"
)

type Project struct {
	gorm.Model
	id uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	projectName string `gorm:"size:100;not null;unique" json:"projectName"`
	projectDescription string `gorm:"size:255;not null" json:"projectDescription"`
	projectLocation string `gorm:"size:255;not null" json:"projectLocation"`
	startDate *time.Time `json:"startDate"`
	endDate *time.Time `json:"endDate"`
	goal int `gorm:"default:0" json:"goal"`
	pledged int `gorm:"default:0" json:"pledged"`
	investors int `gorm:"default:0" json:"investors"`
	// projectStatusId
	// organizationId 
	// userId
}