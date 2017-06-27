package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Cache struct {
	db *gorm.DB
}

func New() *Cache {
	db, err := gorm.Open("sqlite3", "cache.db")
	if err != nil {
		panic("Cannot create cache database")
	}
	cache := &Cache {
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