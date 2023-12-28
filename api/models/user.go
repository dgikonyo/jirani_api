package models

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"html"
	"log"
	"strings"
	"time"
)

type User struct {
	gorm.Model
	ID                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	FirstName         string    `gorm:"size:100;" json:"first_name"`
	LastName          string    `gorm:"size:100;" json:"last_name"`
	Email             string    `gorm:"size:100;" json:"email"`
	Password          string    `gorm:"size:100;" json:"password"`
	ProjectsSupported uint      `gorm:"default:0;" json:"projects_supported"`
	TotalAmount       uint      `gorm:"default:0;" json:"total_amount"`
	CountryID         uint      `gorm:"not null;" json:"country_id"`
	RoleID            uint      `gorm:"not null;" json:"role_id"`
}

func Hash(Password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, Password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(Password))
}

func (user *User) BeforeSave() error {
	hashedPassword, err := Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func (user *User) Prepare() {
	user.ID = uuid.New()
	user.FirstName = user.FirstName
	user.LastName = user.LastName
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	user.ProjectsSupported = user.ProjectsSupported
	user.TotalAmount = user.TotalAmount
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
			"password":          user.Password,
			"firstName":         user.FirstName,
			"lastName":          user.LastName,
			"projectsSupported": user.ProjectsSupported,
			"totalAmount":       user.TotalAmount,
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
