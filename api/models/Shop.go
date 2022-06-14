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
	"golang.org/x/crypto/bcrypt"
)

type Shop struct {
	ID          uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name        string    `gorm:"size:255;not null;unique" json:"name"`
	Email       string    `gorm:"size:100;not null;unique" json:"email"`
	PhoneNumber string    `gorm:"size:50;unique" json:"phone_number"`
	Password    string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (shop *Shop) BeforeSave() error {
	hashedPassword, err := Hash(shop.Password)
	if err != nil {
		return err
	}
	shop.Password = string(hashedPassword)
	return nil
}

func (shop *Shop) Prepare() {
	shop.ID = 0
	shop.Name = html.EscapeString(strings.TrimSpace(shop.Name))
	shop.PhoneNumber = html.EscapeString(strings.TrimSpace(shop.PhoneNumber))
	shop.Email = html.EscapeString(strings.TrimSpace(shop.Email))
	shop.CreatedAt = time.Now()
	shop.UpdatedAt = time.Now()
}

func (shop *Shop) Validate(action string) error {
	fmt.Println("Validating Shop with action:", action)
	switch strings.ToLower(action) {
	case "update":
		if shop.Name == "" {
			return errors.New("Required Name")
		}
		if shop.Password == "" {
			return errors.New("Required Password")
		}
		if shop.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(shop.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if shop.Password == "" {
			return errors.New("Required Password")
		}
		if shop.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(shop.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if shop.Name == "" {
			return errors.New("Required Name")
		}
		if shop.Password == "" {
			return errors.New("Required Password")
		}
		if shop.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(shop.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

func (shop *Shop) SaveShop(db *gorm.DB) (*Shop, error) {

	var err error
	err = db.Debug().Create(&shop).Error
	if err != nil {
		return &Shop{}, err
	}
	return shop, nil
}

func (shop *Shop) FindAllShops(db *gorm.DB) (*[]Shop, error) {
	var err error
	shops := []Shop{}
	// db.Select("id", "name", "email", "phone_number").Find(&shops)
	println(shops)
	err = db.Debug().Model(&Shop{}).Select([]string{"name", "id", "email", "phone_number"}).Limit(100).Find(&shops).Error
	if err != nil {
		return &[]Shop{}, err
	}
	return &shops, err
}

func (shop *Shop) FindShopByID(db *gorm.DB, uid uint32) (*Shop, error) {
	var err error
	err = db.Debug().Model(Shop{}).Where("id = ?", uid).Take(&shop).Error
	if err != nil {
		return &Shop{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Shop{}, errors.New("Shop Not Found")
	}
	shop.Password = ""
	return shop, err
}

func (shop *Shop) UpdateAShop(db *gorm.DB, uid uint32) (*Shop, error) {

	// To hash the password
	err := shop.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&Shop{}).Where("id = ?", uid).Take(&Shop{}).UpdateColumns(
		map[string]interface{}{
			"password":     shop.Password,
			"name":         shop.Name,
			"phone_number": shop.PhoneNumber,
			"email":        shop.Email,
			"update_at":    time.Now(),
		},
	)
	if db.Error != nil {
		return &Shop{}, db.Error
	}
	// This is the display the updated shop
	err = db.Debug().Model(&Shop{}).Where("id = ?", uid).Take(&shop).Error
	if err != nil {
		return &Shop{}, err
	}
	return shop, nil
}

func (shop *Shop) DeleteAShop(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Shop{}).Where("id = ?", uid).Take(&Shop{}).Delete(&Shop{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
