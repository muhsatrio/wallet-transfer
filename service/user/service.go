package user

//go:generate mockgen -destination=../../../mocks/user/usecase_mock.go -package=user_mock -source=initiator.go

import (
	"financial-planner-be/platform/mysql"
)

// Usecase contains the function of business logic of domain user
type Usecase interface {
}

type Service struct {
	UserRepo mysql.UserAdapter
}

var _ Usecase = Service{}
