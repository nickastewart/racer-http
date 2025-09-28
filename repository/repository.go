package repository

import(
	"context"
	"racer_http/sqlite/entities"
)

type UserRepository interface {
	GetUserById(ctx context.Context, id int64) (entities.GetUserByIdRow, error)
	GetUserByEmail(ctx context.Context, email string) (entities.GetUserByEmailRow, error)
	GetUserByEmailForLogin(ctx context.Context, email string) (entities.GetUserByEmailForLoginRow, error)
	CreateUser(ctx context.Context, createUserParams entities.CreateUserParams) (entities.CreateUserRow, error)
}
