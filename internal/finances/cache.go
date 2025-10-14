package finances

import "gorm.io/gorm"

type CacheRepository interface {
	SetMonth(year, month int)
	GetMonth(year, month int)
	InvalidateMonth(year, month int)
}

type cacheRepository struct {
	db *gorm.DB
}

func NewFinanceCacheRepository(db *gorm.DB) CacheRepository {
	return &cacheRepository{db: db}
}

func (c *cacheRepository) SetMonth(year, month int) {
}

func (c *cacheRepository) GetMonth(year, month int) {
}

func (c *cacheRepository) InvalidateMonth(year, month int) {
	c.db.Where("year = ? AND month = ?", year, month).Delete(&MonthCache{})
}
