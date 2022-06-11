package models

import (
	"errors"
	"fmt"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID          uint32    `gorm:"primary_key;auto_increment" json:"id"`
	FirstName   string    `gorm:"size:255;not null;" json:"first_name"`
	LastName    string    `gorm:"size:255;not null;" json:"last_name"`
	Email       string    `gorm:"size:100;not null;unique" json:"email"`
	PhoneNumber string    `gorm:"size:50;unique" json:"phone_number"`
	Password    string    `gorm:"size:100;not null;" json:"-"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
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
	user.ID = 0
	user.FirstName = html.EscapeString(strings.TrimSpace(user.FirstName))
	user.LastName = html.EscapeString(strings.TrimSpace(user.LastName))
	user.PhoneNumber = html.EscapeString(strings.TrimSpace(user.PhoneNumber))
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
}

func (user *User) Validate(action string) error {
	fmt.Println("Validating User with action:", action)
	switch strings.ToLower(action) {
	case "update":
		if user.FirstName == "" {
			return errors.New("Required First Name")
		}
		if user.LastName == "" {
			return errors.New("Required Last Name")
		}
		if user.Password == "" {
			return errors.New("Required Password")
		}
		if user.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if user.Password == "" {
			return errors.New("Required Password")
		}
		if user.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if user.FirstName == "" {
			return errors.New("Required First Name")
		}
		if user.LastName == "" {
			return errors.New("Required Last Name")
		}
		if user.Password == "" {
			return errors.New("Required Password")
		}
		if user.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid Email")
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

	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (user *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return user, err
}

func (user *User) UpdateAUser(db *gorm.DB, uid uint32) (*User, error) {

	// To hash the password
	err := user.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":     user.Password,
			"first_name":   user.FirstName,
			"last_name":    user.LastName,
			"phone_number": user.PhoneNumber,
			"email":        user.Email,
			"update_at":    time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	// This is the display the updated User
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
