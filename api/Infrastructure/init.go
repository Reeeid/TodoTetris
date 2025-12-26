package infrastructure

import (
	usecase "github.com/Reeeid/TodoTetris/UseCase"
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
	UserUseCase *usecase.UserUseCase
	TodoUseCase *usecase.TodoUseCase
	GameUseCase *usecase.GameUseCase
)

func init() {
	DBprovider := NewSupabaseDB()
	db := DBprovider.GetDB()

	//依存性注入で各サービスの立ち上げ
	TodoRepo := NewTodoRepo(db)
	UserRepo := NewUserRepo(db)
	GameRepo := NewGameRepo(db)
	UserUseCase = usecase.NewUserUseCase(UserRepo)
	TodoUseCase = usecase.NewTodoUseCase(TodoRepo)
	GameUseCase = usecase.NewGameUseCase(GameRepo)
}
