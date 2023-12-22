package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	id                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	firstName         string
	lastName          string
	email             string
	password          string
	projectsSupported uint `gorm:"default:0;"`
	totalAmount       uint `gorm:"default:0;"`
	CountryID         uint `gorm:"not null;"`
	RoleID            uint `gorm:"not null;"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (user *User) BeforeSave() error {
	hashedPassword, err := Hash(user.password)
	if err != nil {
		return err
	}
	user.password = string(hashedPassword)
	return nil
}

func (user *User) Prepare() {
	user.id = uuid.New()
	user.firstName = user.firstName
	user.lastName = user.lastName
	user.email = html.EscapeString(strings.TrimSpace(user.email))
	user.projectsSupported = user.projectsSupported
	user.totalAmount = user.totalAmount
	user.CountryID = user.CountryID
	user.RoleID = user.RoleID
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
}

func (user *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&user).Error
	if err != nil {
		return &User{}, err
	}

	return user, nil
}

func (user *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Preload("Countries").Limit(100).Find(&users).Error

	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (user *User) FindUserById(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return user, err
}

func (user *User) UpdateAUser(db *gorm.DB, uid uint32) (*User, error) {
	err := user.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":          user.password,
			"firstName":         user.firstName,
			"lastName":          user.lastName,
			"projectsSupported": user.projectsSupported,
			"totalAmount":       user.totalAmount,
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}

	//lets display the updated user
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Delete(&user)

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
