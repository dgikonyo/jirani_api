package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

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
	roleId            uint      `gorm:"not null;"json:"roleId"`
	country           Country   `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL;"`
	Countries         []Country
	role              Role `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL;"`
	Roles             []Role
	CreatedAt         time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt         time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt         time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"deleted_at"`
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
	user.countryId = user.countryId
	user.roleId = user.roleId
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
}

func (user *User) Validate(action string) error {
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
		if user.roleId == 0 {
			return errors.New("Required Role")
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
