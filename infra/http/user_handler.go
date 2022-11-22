package http

import (
	"clean-arch/usecase/user"
)

type UserHandler struct {
	userUC user.UseCase
}

func NewUserHandler(uc user.UseCase) UserHandler {
	return UserHandler{
		userUC: uc,
	}
}

// Omitted.
