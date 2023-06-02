package models

import "gorm.io/gorm"

type CollectedPage struct {
	gorm.Model
	Path     string `gorm:"size:255;not null"`
	XssHitID uint
	Content  []byte `gorm:"serializer:json" json:"-"`
}
