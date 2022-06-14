package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Product struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `gorm:"size:255;not null;" json:"title,omitempty" bson:",omitempty"`
	Description string    `gorm:"size:255;not null;" json:"description,omitempty" bson:",omitempty"`
	Shop        Shop      `json:"-"`
	ShopID      uint32    `gorm:"not null" json:"shop_id,omitempty" bson:",omitempty"`
	ImageUrl    string    `gorm:"size: 255;" json:"image_url,omitempty" bson:",omitempty"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Product) Prepare() {
	p.ID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Description = html.EscapeString(strings.TrimSpace(p.Description))
	p.Shop = Shop{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Product) Validate(params ...int) error {
	if params == nil {
		if p.Title == "" {
			return errors.New("Required Title")
		}
		if p.Description == "" {
			return errors.New("Required Description")
		}
		if p.ShopID < 1 {
			return errors.New("Required Shop")
		}
		return nil
	}
	return nil

}

func (p *Product) SaveProduct(db *gorm.DB) (*Product, error) {
	var err error
	err = db.Debug().Model(&Product{}).Create(&p).Error
	if err != nil {
		return &Product{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&Shop{}).Where("id = ?", p.ShopID).Take(&p.Shop).Error
		if err != nil {
			return &Product{}, err
		}
	}
	return p, nil
}

func (p *Product) FindAllProducts(db *gorm.DB) (*[]Product, error) {
	var err error
	products := []Product{}
	err = db.Debug().Model(&Product{}).Limit(100).Find(&products).Error
	if err != nil {
		return &[]Product{}, err
	}
	if len(products) > 0 {
		for i, _ := range products {
			err := db.Debug().Model(&Shop{}).Where("id = ?", products[i].ShopID).Take(&products[i].Shop).Error
			if err != nil {
				return &[]Product{}, err
			}
		}
	}
	return &products, nil
}

func (p *Product) FindProductByID(db *gorm.DB, pid uint64) (*Product, error) {
	var err error
	err = db.Debug().Model(&Product{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Product{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&Shop{}).Where("id = ?", p.ShopID).Take(&p.Shop).Error
		if err != nil {
			return &Product{}, err
		}
	}
	return p, nil
}

func (p *Product) UpdateAProduct(db *gorm.DB) (*Product, error) {

	var err error

	err = db.Debug().Model(&Product{}).Where("id = ?", p.ID).Updates(Product{Title: p.Title, Description: p.Description, UpdatedAt: time.Now(), ImageUrl: p.ImageUrl}).Error
	if err != nil {
		return &Product{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&Shop{}).Where("id = ?", p.ShopID).Take(&p.Shop).Error
		if err != nil {
			return &Product{}, err
		}
	}
	return p, nil
}

func (p *Product) DeleteAProduct(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Product{}).Where("id = ? and shop_id = ?", pid, uid).Take(&Product{}).Delete(&Product{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Product not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
