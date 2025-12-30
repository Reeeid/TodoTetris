package init

import (
	infrastructure "github.com/Reeeid/TodoTetris/Infrastructure"
	usecase "github.com/Reeeid/TodoTetris/UseCase"
)

var (
	TodoUsecase *usecase.TodoUseCase
	GameUsecase *usecase.GameUseCase
	UUIDUsecase *usecase.UUIDUseCase
	UserUsecase *usecase.UserUseCase
)

func init() {
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
