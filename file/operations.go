package file

import (
	"os"
	"path/filepath"
	"log"
	"github.com/antonyho/go-dupefinder/database"
	"syscall"
	"time"
)

func Cache(cache *database.Cache) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Panicf("Problem analysing file. Error: %v", err)
			return nil
		}
		fileStat := info.Sys().(*syscall.Stat_t)
		if !info.IsDir() && (info.Size() > 0) {
			var hash string		// TODO Digest the file
			f := database.File{
				Path: path,
				Hash: hash,
				Size: info.Size(),
				CreationTime: time.Unix(int64(fileStat.Ctimespec.Sec), int64(fileStat.Ctimespec.Nsec)),
				ModificationTime: info.ModTime(),
			}
			cache.Add(f)
		}

		return nil
	}
}