package infrastructure

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewSupabaseDB() *SupabaseDBProvider {
	return &SupabaseDBProvider{}
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

var (
	TodoRepo *TodoRepoImpl
	UserRepo *UserRepoImpl
	GameRepo *GameRepoImpl
)

func init() {
	DBprovider := NewSupabaseDB()
	db := DBprovider.GetDB()
	TodoRepo = NewTodoRepo(db)
	UserRepo = NewUserRepo(db)
	GameRepo = NewGameRepo(db)

}
