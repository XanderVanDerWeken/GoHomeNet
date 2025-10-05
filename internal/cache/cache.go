package cache

import (
	"log"
	"sync"

	"github.com/xandervanderweken/GoHomeNet/internal/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	cache *gorm.DB
	once  sync.Once
)

func ConnectCache() *gorm.DB {
	once.Do(func() {
		cfg := config.AppConfig.Cache

		db, err := gorm.Open(sqlite.Open(cfg.Path), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed tp connect to cache db: %v", err)
		}

		cache = db
	})

	return cache
}
