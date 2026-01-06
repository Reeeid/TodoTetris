package infrastructure

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewSupabaseDB() *SupabaseDBProvider {
	return &SupabaseDBProvider{}
}

type SupabaseDBProvider struct{}

func (p *SupabaseDBProvider) GetDB() *gorm.DB {
	dsn := os.Getenv("SUPABASE_URL")

	// Disable prepared statements for Supabase Transaction Pooler compatibility
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: false,
	})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get sqlDB: " + err.Error())
	}

	sqlDB.SetMaxOpenConns(1)
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetConnMaxLifetime(60 * 60)

	return db
}
