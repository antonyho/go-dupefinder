package database

import (
	"crypto/sha1"
	"fmt"
	"github.com/antonyho/go-dupefinder/file"
	hash2 "hash"
	"io"
	"log"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

func Store(cache *Cache) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("\nProblem analysing file %s\nError: %v\n", path, err)
			return nil
		}
		fileStat := info.Sys().(*syscall.Stat_t)
		if !info.IsDir() && (info.Size() > 0) {
			var (
				checksum string
				hash     hash2.Hash
			)
			if fp, err := os.Open(path); err != nil {
				log.Printf("\nUnable to open %s\n", path)
				return nil
			} else {
				defer fp.Close()
				hash = sha1.New()
				if _, err := io.Copy(hash, fp); err != nil {
					log.Printf("\nError hashing file checksum %s\n", path)
					return nil
				}
				checksum = fmt.Sprint("%x", hash.Sum(nil))
			}
			f := file.Info{
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
