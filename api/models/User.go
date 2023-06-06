package models

import (
	"errors"
	"html"
	"log"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	id                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	firstName         string    `gorm:"size:255;not null;" json:"firstName"`
	lastName          string    `gorm:"size:255;not null;" json:"lastName"`
	email             string    `gorm:"size:100;not null;unique" json:"email"`
	password          string    `gorm:"size:100;not null;" json:"password"`
	projectsSupported uint      `gorm:"defualt:0" json:"projectsSupported"`
	totalAmount       uint      `gorm:"defualt:0" json:"totalAmount"`
	countryId         uint      `gorm:"not null;"json:"countryId"`
	country           Country   `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL;"`
	Countries         []Country
}

func hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (user *User) beforeSave() error {
	hashedPassword, err := hash(user.password)
	if err != nil {
		return err
	}
	user.password = string(hashedPassword)
	return nil
}

func (user *User) prepare() {
	user.id = uuid.New()
	user.firstName = user.firstName
	user.lastName = user.lastName
	user.email = html.EscapeString(strings.TrimSpace(user.email))
	user.projectsSupported = user.projectsSupported
	user.totalAmount = user.totalAmount
}

func (user *User) validate(action string) error {
	switch strings.ToLower(action) {
	case "registration":
		if user.firstName == "" {
			return errors.New("Required First Name")
		}
		if user.lastName == "" {
			return errors.New("Required Last Name")
		}
		if user.email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(user.email); err != nil {
			return errors.New("Invalid Email")
		}
		if user.countryId == 0 {
			return errors.New("Required Country")
		}

		return nil
	case "login":
		if user.email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(user.email); err != nil {
			return errors.New("Invalid Email")
		}
		if user.password == "" {
			return errors.New("Required Password")
		}

		return nil
	default:
		if user.firstName == "" {
			return errors.New("Required First Name")
		}
		if user.lastName == "" {
			return errors.New("Required Last Name")
		}
		if user.email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(user.email); err != nil {
			return errors.New("Invalid Email")
		}
		if user.countryId == 0 {
			return errors.New("Required Country")
		}
		return nil
	}
}

func (user *User) saveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&user).Error
	if err != nil {
		return &User{}, err
	}

	return user, nil
}

func (user *User) findAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Preload("Countries").Limit(100).Find(&users).Error

	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (user *User) findUserById(db *gorm.DB, uid uint32) (*User, error) {
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

func (user *User) updateAUser(db *gorm.DB, uid uint32) (*User, error) {
	err := user.beforeSave()
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
