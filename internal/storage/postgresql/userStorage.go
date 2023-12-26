package postgresql

import (
	"PartyRoom.API/internal/domain"
	"github.com/google/uuid"
)

func (s *Storage) CreateUser(user domain.User) error {
	var role domain.Role
	result := s.db.Where("name = ?", domain.RoleNameUser).First(&role)
	if result.Error != nil {
		return result.Error
	}
	user.Roles = []*domain.Role{&role}
	result = s.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Storage) GetUserById(uuid uuid.UUID) (*domain.User, error) {
	var user *domain.User
	result := s.db.Preload("Roles").Where("id = ?", uuid).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (s *Storage) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	result := s.db.Preload("Roles").Where("email =?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
