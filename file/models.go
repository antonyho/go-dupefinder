package file

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Info struct {
	gorm.Model
	Path             string    `csv:"file path"`
	Hash             string    `gorm:"index" csv:"checksum"`
	Size             int64     `gorm:"index" csv:"size"`
	Type             string    `csv:"-"`
	CreationTime     time.Time `csv:"creation time"`
	ModificationTime time.Time `csv:"modification time"`
}
func (Info) TableName() string {
	return "fileinfo"
}

type ChecksumResult struct {
	Hash  string
	Total int64
}

type Group struct {
	Files    []Info
	Checksum string
}

type Files []Info

func (s Files) Len() int {
	return len(s)
}
func (s Files) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type BySize struct {
	Files
}

func (s BySize) Less(i, j int) bool {
	return s.Files[i].Size < s.Files[j].Size
}

type ByCreationTime struct {
	Files
}

func (s ByCreationTime) Less(i, j int) bool {
	return s.Files[i].CreationTime.Before(s.Files[j].CreationTime)
}

type ByModificationTime struct {
	Files
}

func (s ByModificationTime) Less(i, j int) bool {
	return s.Files[i].ModificationTime.Before(s.Files[j].ModificationTime)
}
