package database

import (
	"github.com/antonyho/go-dupefinder/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
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
	if err = cache.Initialise(); err != nil {
		panic("Cannot create table")
	}

	return cache
}

func (c Cache) Initialise() error {
	// Create table or update table
	if err := c.db.AutoMigrate(&file.Info{}).Error; err != nil {
		return err
	}
	return nil
}

func (c Cache) Add(f file.Info) error {
	if err := c.db.Create(&f).Error; err != nil {
		return err
	}
	return nil
}

func (c Cache) Close() {
	c.db.Close()
	// TODO Delete the db file if it is not on memory
}

func (c Cache) ListDuplicated() ([]file.Group, error) {
	tableName := (file.Info{}).TableName()
	var groups []file.Group
	results, err := c.db.Table(tableName).Order("size desc").Select("hash, count(1) as total").Group("hash").Having("count(1) > ?", 1).Rows()
	if err != nil {
		return nil, err
	}
	for results.Next() {
		var (
			checksum string
			total int
			group file.Group
		)
		group.Files = make([]file.Info, 0)
		if err = results.Scan(&checksum, &total); err != nil {
			return nil, err
		}
		var cnt int
		c.db.Find(&group.Files, file.Info{Hash: checksum}).Count(&cnt)
		log.Printf("Count: %v", cnt)
		group.Checksum = group.Files[0].Hash
		groups = append(groups, group)
	}

	return groups, nil
}
