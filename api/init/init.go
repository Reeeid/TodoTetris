package init

import (
	"fmt"

	infrastructure "github.com/Reeeid/TodoTetris/api/Infrastructure"
	usecase "github.com/Reeeid/TodoTetris/api/UseCase"
)

var (
	TodoUsecase *usecase.TodoUseCase
	GameUsecase *usecase.GameUseCase
	UUIDUsecase *usecase.UUIDUseCase
	UserUsecase *usecase.UserUseCase
)

func init() {
	fmt.Println("DEBUG: api/init/init.go initializing...")
	driver := infrastructure.NewSupabaseDB()
	db := driver.GetDB()
	TodoRepo := infrastructure.NewTodoRepo(db)
	TodoUsecase = usecase.NewTodoUseCase(TodoRepo)
	GameRepo := infrastructure.NewGameRepo(db)
	GameUsecase = usecase.NewGameUseCase(GameRepo)
	UserRepo := infrastructure.NewUserRepo(db)
	UserUsecase = usecase.NewUserUseCase(UserRepo)
	UUIDUsecase = usecase.NewUUIDUseCase()
}
