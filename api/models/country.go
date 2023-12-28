package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"strings"
)

type Country struct {
	gorm.Model
	Id          uint   `gorm:"primary_key;auto_increment" json:"id"`
	CountryName string `gorm:"size:255;not null;unique" json:"CountryName"`
	CountryCode uint64 `gorm:"not null;unique" json:"CountryCode"`
	Users       []User `gorm:"foreignKey:CountryID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (country *Country) prepare() {
	country.Id = 0
	country.CountryName = country.CountryName
	country.CountryCode = country.CountryCode
}

func (country *Country) validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if country.CountryName == "" {
			return errors.New("Required Country name")
		}
		if country.CountryCode == 0 {
			return errors.New("Required Country Code")
		}
		return nil
	default:
		if country.CountryName == "" {
			return errors.New("Required Country name")
		}
		if country.CountryCode == 0 {
			return errors.New("Required Country Code")
		}
		return nil
	}
}

func (country *Country) saveCountry(db *gorm.DB) (*Country, error) {
	var err error
	err = db.Debug().Model(&Country{}).Create(&country).Error

	if err != nil {
		return &Country{}, err
	}
	return country, nil
}

func (country *Country) findAllCountries(db *gorm.DB) (*[]Country, error) {
	var err error
	countries := []Country{}
	err = db.Debug().Model(&Country{}).Limit(100).Find(&countries).Error

	if err != nil {
		return &[]Country{}, err
	}
	return &countries, err
}

func (country *Country) updateACountry(db *gorm.DB, uid uint32) (*Country, error) {
	db = db.Debug().Model(&Country{}).Where("id = ?", uid).Take(&Country{}).UpdateColumns(
		map[string]interface{}{
			"CountryName": country.CountryName,
			"CountryCode": country.CountryCode,
		},
	)
	if db.Error != nil {
		return &Country{}, db.Error
	}

	//display updated user
	err := db.Debug().Model(&Country{}).Where("id = ?", uid).Take(&country).Error
	if err != nil {
		return &Country{}, err
	}
	return country, nil
}

func (country *Country) DeleteACountry(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Country{}).Where("id = ?", uid).Take(&Country{}).Delete(&Country{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
