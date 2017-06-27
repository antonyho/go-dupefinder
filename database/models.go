package database

import (
	"time"
	"github.com/jinzhu/gorm"
)

type File struct {
	gorm.Model
	Path             string
	Hash             string    `gorm:"index"`
	Size             int64     `gorm:"index"`
	Type             string
	CreationTime     time.Time
	ModificationTime time.Time
}
