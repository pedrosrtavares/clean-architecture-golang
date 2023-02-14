package service

import (
	"clean-architecture-golang/internal/entity"
	"clean-architecture-golang/internal/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func (s *UserService) GetByID(id string) (interface{}, error) {
	return s.userRepo.FindByID(id)
}

func (s *UserService) GetByEmail(email string) (interface{}, error) {
	return s.userRepo.FindByEmail(email)
}

func (s *UserService) Create(data entity.User) (interface{}, error) {
	return s.userRepo.CreateUser(&data), nil
}

func (s *UserService) Update(data entity.User) (interface{}, error) {
	return s.userRepo.UpdateUser(&data), nil
}

func (s *UserService) Delete(id string) error {
	return s.userRepo.DeleteUser(id)
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}
