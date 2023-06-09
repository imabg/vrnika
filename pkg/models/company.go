package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Company struct {
	ID                  uint           `gorm:"primaryKey;autoIncrement;unique" json:"id"`
	Name                string         `gorm:"not null" json:"name" binding:"required"`
	URL                 string         `gorm:"unique; not null" json:"url" binding:"required"`
	Address             datatypes.JSON `json:"address,omitempty"`
	SocialMedia         datatypes.JSON `json:"socialMedia,omitempty"`
	Description         string         `json:"description"`
	IsRemote            bool           `json:"isRemote"`
	IsJobPostingDisable bool           `gorm:"default:false" json:"isJobPostingDisable"`
	Users               []User         `json:"users,omitempty"`
	CreatedAt           int64          `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt           int64          `gorm:"autoUpdateTime" json:"updatedAt"`
}

// Create
func (c *Company) Create(db *gorm.DB) error {
	err := db.Create(c).Error
	if err != nil {
		return err
	}
	return nil
}

// Get details
func (c *Company) GetDetails(db *gorm.DB) error {
	err := db.Find(&c).Error
	if err != nil {
		return err
	}
	return nil
}
