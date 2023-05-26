package models

import (
	"gorm.io/gorm"
)

type Company struct {
	ID   uint   `gorm:"primaryKey;autoIncrement;unique" json:"id"`
	Name string `gorm:"not null" json:"name" binding:"required"`
	URL  string `gorm:"unique; not null" json:"url" binding:"required"`
	// Address             Address       `json:"omitempty"`
	// SocialAccount       SocialAccount `json:"omitempty"`
	// Images              []Images      `json:"omitempty"`
	Description         string `json:"description"`
	IsRemote            bool   `json:"isRemote"`
	IsJobPostingDisable bool   `gorm:"default:false" json:"isJobPostingDisable"`
	CreatedAt           int64  `gorm:"autoCreateTime"`
	UpdatedAt           int64  `gorm:"autoUpdateTime"`
}

type Images struct {
	Path     string `json:"path"`
	Alt      string `json:"alt"`
	Priority uint   `json:"priority"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

type SocialAccount struct {
	Linkedin string `json:"linkedin"`
	Twitter  string `json:"twitter"`
	Youtube  string `json:"youtube"`
}

// Create
func (c *Company) Create(db *gorm.DB) error {
	err := db.Create(c).Error
	if err != nil {
		return err
	}
	return nil
}
