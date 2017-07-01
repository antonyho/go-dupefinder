package file

import "github.com/antonyho/go-dupefinder/database"

type Group struct {
	Files    []database.File
	Checksum string
}

type Files []database.File

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
