package repository

import (
	"context"
	"racer_http/sqlite/entities"
)

type UserRepository interface {
	GetUserById(ctx context.Context, id int64) (entities.GetUserByIdRow, error)
	GetUserByEmail(ctx context.Context, email string) (entities.GetUserByEmailRow, error)
	GetUserByEmailForLogin(ctx context.Context, email string) (entities.GetUserByEmailForLoginRow, error)
	CreateUser(ctx context.Context, createUserParams entities.CreateUserParams) (entities.CreateUserRow, error)
}

type LocationRepository interface {
	CreateLocation(ctx context.Context, name string) (entities.Location, error)
	GetLocationByName(ctx context.Context, name string) (entities.Location, error)
}

type EventRepository interface {
	CreateEvent(ctx context.Context, arg entities.CreateEventParams) (entities.Event, error)
	GetEventByLocationAndTypeAndDate(ctx context.Context, arg entities.GetEventByLocationAndTypeAndDateParams) (entities.Event, error)
}
