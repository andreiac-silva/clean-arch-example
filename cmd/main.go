package main

import (
	"clean-arch/infra/http"
	"clean-arch/infra/persistence"
	"clean-arch/usecase/book"
	"clean-arch/usecase/user"
)

// Dependency injection to avoid a long main.go file? "Wire" not?

func main() {
	bookRepo := persistence.NewBookRepository()
	userRepo := persistence.NewUserRepository()

	bookUC := book.NewUseCase(bookRepo)
	userUC := user.NewUseCase(userRepo)

	// Just an architectural test. No http server exposed. Maybe in the future :)
	_ = http.NewBookHandler(bookUC)
	_ = http.NewUserHandler(userUC)
}
