package repository

import (
	"clean-architecture-golang/internal/entity"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	FindByID(id string) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	CreateUser(user *entity.User) error
	UpdateUser(user *entity.User) error
	DeleteUser(id string) error
}

type PostgresUserRepositoryImpl struct {
	DB *sql.DB
}

func (p PostgresUserRepositoryImpl) FindByID(id string) (*entity.User, error) {
	var entity entity.User
	err := p.DB.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&entity.ID, &entity.FirstName, &entity.LastName, &entity.Email, &entity.Password, &entity.PhoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with id %s not found", id)
		}
		return nil, fmt.Errorf("error retrieving user: %v", err)
	}
	return &entity, nil
}

func (p PostgresUserRepositoryImpl) FindByEmail(email string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostgresUserRepositoryImpl) CreateUser(user *entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (p PostgresUserRepositoryImpl) UpdateUser(user *entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (p PostgresUserRepositoryImpl) DeleteUser(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewPostgresUserRepository(db *sql.DB) UserRepository {
	return &PostgresUserRepositoryImpl{
		DB: db,
	}
}
