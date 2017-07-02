package database

import (
	"github.com/antonyho/go-dupefinder/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Cache struct {
	db *gorm.DB
}

func New() *Cache {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic("Cannot create cache database")
	}
	cache := &Cache{
		db: db,
	}
	cache.Initialise()

	return cache
}

func (c Cache) Initialise() {
	// Create table or update table
	c.db.AutoMigrate(&file.Info{})
}

func (c Cache) Add(f file.Info) {
	c.db.Create(&f)
}

func (c Cache) Close() {
	c.db.Close()
	// TODO Delete the db file if it is not on memory
}

func (c Cache) ListDuplicated() ([]file.Group, error) {
	var groups []file.Group
	results, err := c.db.Table("files").Order("size desc").Select("hash, count(1) as total").Group("hash").Having("count(1) > ?", 1).Rows()
	if err != nil {
		return nil, err
	}
	for results.Next() {
		var (
			checksumResult file.ChecksumResult
			group          file.Group
		)
		if err = results.Scan(&checksumResult); err != nil {
			return nil, err
		}
		c.db.Where(&file.Info{Hash: checksumResult.Hash}).Find(&(group.Files))
		group.Checksum = group.Files[0].Hash
		groups = append(groups, group)
	}

	return groups, nil
}
