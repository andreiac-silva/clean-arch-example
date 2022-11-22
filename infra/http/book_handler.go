package http

import (
	"clean-arch/usecase/book"
)

type BookHandler struct {
	bookUC book.UseCase
}

func NewBookHandler(uc book.UseCase) BookHandler {
	return BookHandler{
		bookUC: uc,
	}
}

// Omitted.
