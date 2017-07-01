package file

import (
	"crypto/sha1"
	"fmt"
	"github.com/antonyho/go-dupefinder/database"
	hash2 "hash"
	"io"
	"log"
	"os"
	"path/filepath"
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
			var (
				checksum string
				hash     hash2.Hash
			)
			if fp, err := os.Open(path); err != nil {
				log.Panicf("Unable to open %s.\n", path)
				return err
			} else {
				defer fp.Close()
				hash = sha1.New()
				if _, err := io.Copy(hash, fp); err != nil {
					log.Panicf("Error hashing file checksum %s.\n", path)
					return err
				}
				checksum = fmt.Sprint("%x", hash.Sum(nil))
			}
			f := database.File{
				Path:             path,
				Hash:             checksum,
				Size:             info.Size(),
				CreationTime:     time.Unix(int64(fileStat.Ctimespec.Sec), int64(fileStat.Ctimespec.Nsec)),
				ModificationTime: info.ModTime(),
			}
			cache.Add(f)
		}

		return nil
	}
}
