package database

import (
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
	c.db.AutoMigrate(&File{})
}

func (c Cache) Add(f File) {
	c.db.Create(&f)
}

func (c Cache) Close() {
	c.db.Close()
	// TODO Delete the db file if it is not on memory
}
