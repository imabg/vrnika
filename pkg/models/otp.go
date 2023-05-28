package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OTP struct {
	Id        uint      `gorm:"primaryKey;autoIncrement;unique" json:"id"`
	OTP       string    `gorm:"not null" json:"otp"`
	MappingId uuid.UUID `gorm:"not null" json:"mappingId"`
	CreatedAt int64     `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt int64     `gorm:"autoUpdateTime" json:"updatedAt"`
}

func (o *OTP) Save(db *gorm.DB) error {
	err := db.Create(&o).Error

	if err != nil {
		return err
	}
	return nil
}

func (o *OTP) GetLatestOTP(db *gorm.DB) error {
	tx := db.Begin()
	db.Find(&o)
	db.Delete(&o)
	if err := tx.Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
