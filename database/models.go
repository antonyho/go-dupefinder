package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

type File struct {
	gorm.Model
	Path             string
	Hash             string `gorm:"index"`
	Size             int64  `gorm:"index"`
	Type             string
	CreationTime     time.Time
	ModificationTime time.Time
}

type ChecksumResult struct {
	Hash  string
	Total int64
}
