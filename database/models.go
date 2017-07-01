package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

type File struct {
	gorm.Model
	Path             string              `csv:"file path"`
	Hash             string `gorm:"index" csv:"checksum"`
	Size             int64  `gorm:"index" csv:"size"`
	Type             string              `csv:"-"`
	CreationTime     time.Time           `csv:"creation time"`
	ModificationTime time.Time           `csv:"modification time"`
}

type ChecksumResult struct {
	Hash  string
	Total int64
}
