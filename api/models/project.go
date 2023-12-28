package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type Project struct {
	gorm.Model
	ID                 uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	ProjectName        string     `gorm:"size:100;not null;unique" json:"projectName"`
	ProjectDescription string     `gorm:"size:255;not null" json:"projectDescription"`
	ProjectLocation    string     `gorm:"size:255;not null" json:"projectLocation"`
	StartDate          *time.Time `json:"startDate"`
	EndDate            *time.Time `json:"endDate"`
	Goal               int        `gorm:"default:0" json:"goal"`
	Pledged            int        `gorm:"default:0" json:"pledged"`
	Investors          int        `gorm:"default:0" json:"investors"`
	Comments           []Comment  `gorm:"foreignKey:ProjectID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	// projectStatusId
	// organizationId
	// userId
}
