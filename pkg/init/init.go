package init

import (
	infrastructure "github.com/Reeeid/TodoTetris/pkg/Infrastructure"
	usecase "github.com/Reeeid/TodoTetris/pkg/UseCase"
)

var (
	todoUsecase *usecase.TodoUseCase
	gameUsecase *usecase.GameUseCase
	uuidUsecase *usecase.UUIDUseCase
	userUsecase *usecase.UserUseCase
)

func GetTodoUsecase() *usecase.TodoUseCase {
	if todoUsecase == nil {
		driver := infrastructure.NewSupabaseDB()
		db := driver.GetDB()
		repo := infrastructure.NewTodoRepo(db)
		todoUsecase = usecase.NewTodoUseCase(repo)
	}
	return todoUsecase
}

func GetGameUsecase() *usecase.GameUseCase {
	if gameUsecase == nil {
		driver := infrastructure.NewSupabaseDB()
		db := driver.GetDB()
		repo := infrastructure.NewGameRepo(db)
		gameUsecase = usecase.NewGameUseCase(repo)
	}
	return gameUsecase
}

func GetUserUsecase() *usecase.UserUseCase {
	if userUsecase == nil {
		driver := infrastructure.NewSupabaseDB()
		db := driver.GetDB()
		repo := infrastructure.NewUserRepo(db)
		userUsecase = usecase.NewUserUseCase(repo)
	}
	return userUsecase
}

func GetUUIDUsecase() *usecase.UUIDUseCase {
	if uuidUsecase == nil {
		uuidUsecase = usecase.NewUUIDUseCase()
	}
	return uuidUsecase
}
