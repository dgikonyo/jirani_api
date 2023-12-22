package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

type Role struct {
	gorm.Model
	id    uint   `gorm:"primary_key""`
	name  string `gorm:"size:255;not null;"`
	Users []User `gorm:"foreignKey:RoleID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (role *Role) Prepare() {
	role.id = role.id
	role.name = role.name
	role.CreatedAt = time.Now()
	role.UpdatedAt = time.Now()
}

func (role *Role) Validate(action string) error {
	switch strings.ToLower(action) {
	case "creation":
		if role.name == "" {
			return errors.New("Required Role Name")
		}

		return nil
	case "update":
		if role.name == "" {
			return errors.New("Required Role Name")
		}

		return nil

	default:
		if role.name == "" {
			return errors.New("Required Role Name")
		}

		return nil
	}
}

func (role *Role) SaveRole(db *gorm.DB) (*Role, error) {
	var err error
	err = db.Debug().Create(&role).Error

	if err != nil {
		return &Role{}, err
	}
	return role, nil
}

func (role *Role) UpdateARole(db *gorm.DB, uid uint32) (*Role, error) {
	db = db.Debug().Model(&Role{}).Where("id = ?", uid).Take(&Role{}).UpdateColumns(
		map[string]interface{}{
			"name": role.name,
		},
	)

	if db.Error != nil {
		return &Role{}, db.Error
	}

	//lets display the updated user
	err := db.Debug().Model(&Role{}).Where("id = ?", uid).Take(&role).Error
	if err != nil {
		return &Role{}, err
	}
	return role, nil
}
