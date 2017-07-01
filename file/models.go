package file

import "github.com/antonyho/go-dupefinder/database"

type Group struct {
	Files    []database.File
	Checksum string
}
