package infrastructure

import (
	"fmt"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewSupabaseDB() *SupabaseDBProvider {
	return &SupabaseDBProvider{}
}

type SupabaseDBProvider struct{}

func (p *SupabaseDBProvider) GetDB() *gorm.DB {
	// DEBUG: Print all env keys
	fmt.Println("DEBUG: Dumping Environment Keys:")
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println("KEY:", pair[0])
	}

	dsn := os.Getenv("DB_PATH")
	dsn = strings.TrimSpace(dsn)
	fmt.Printf("DEBUG: DB_PATH Length: %d\n", len(dsn))
	if len(dsn) > 15 {
		fmt.Printf("DEBUG: DB_PATH Prefix: %s...\n", dsn[:15])
	}

	if dsn == "" {
		panic("CRITICAL ERROR: DB_PATH environment variable is not set (or empty). Please add it in Vercel Settings.")
	}

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

	// Limit connections to avoid "MaxClientsInSessionMode" error on Supabase
	// especially for Vercel functions which scale out.
	sqlDB.SetMaxOpenConns(1)
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetConnMaxLifetime(60 * 60) // 1 Hour (or shorter if needed)

	return db
}
