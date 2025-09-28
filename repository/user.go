package repository

import (
	"context"
	"log"
	"racer_http/sqlite/entities"
)

type UserRepositorySqlite struct {
	queries *entities.Queries
}

func NewUserRepository(queries *entities.Queries) *UserRepositorySqlite {
	return &UserRepositorySqlite{
		queries: queries,
	}
}

func (r *UserRepositorySqlite) GetUserById(ctx context.Context, id int64) (entities.GetUserByIdRow, error) {
	user, err := r.queries.GetUserById(ctx, id)

	if err != nil {
		log.Fatal(err)
	}
	return user, err
}

func (r *UserRepositorySqlite) CreateUser(ctx context.Context,
	userParams entities.CreateUserParams) (entities.CreateUserRow, error) {

	savedUser, err := r.queries.CreateUser(ctx, userParams)
	if err != nil {
		log.Fatal(err)
	}

	return savedUser, err
}

func (r *UserRepositorySqlite) GetUserByEmail(ctx context.Context, email string) (entities.GetUserByEmailRow, error) {
 	return r.queries.GetUserByEmail(ctx, email)
}

func (r *UserRepositorySqlite) GetUserByEmailForLogin(ctx context.Context, email string) (entities.GetUserByEmailForLoginRow, error) {
 	return r.queries.GetUserByEmailForLogin(ctx, email)
}
