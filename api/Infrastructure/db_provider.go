package infrastructure

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBprovider interface {
	GetDB() *gorm.DB
}

type SupabaseDBProvider struct{}

func (p *SupabaseDBProvider) GetDB() *gorm.DB {
	dsn := "link"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
