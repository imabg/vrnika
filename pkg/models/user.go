package models

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID      `gorm:"unique, not null" json:"id"`
	FirstName   string         `gorm:"not null" json:"firstName"`
	MiddleName  string         `json:"middleName,omitempty"`
	LastName    string         `gorm:"not null" json:"lastName"`
	Address     datatypes.JSON `json:"address,omitempty"`
	Avatar      string         `json:"avatar,omitempty"`
	RoleID      uint           `json:"roleId,omitempty"`
	IsVerified  bool           `json:"isVerified"`
	IsActive    bool           `gorm:"default:true" json:"isActive"`
	Department  string         `gorm:"not null" json:"department"`
	Designation string         `gorm:"not null" json:"designation"`
	Type        string         `gorm:"not null" json:"type"`
	Timezone    string         `gorm:"not null" json:"timezone"`
	CompanyID   uint           `json:"companyId"`
	CreatedAt   int64          `gorm:"autoCreateTime"`
	UpdatedAt   int64          `gorm:"autoUpdateTime"`
}

// Hook - BeforeCreate
func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}

// Create
func (u *User) Create(db *gorm.DB) error {
	err := db.Create(u).Error
	if err != nil {
		return err
	}
	return nil
}
