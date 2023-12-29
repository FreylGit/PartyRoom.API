package postgresql

import (
	"PartyRoom.API/internal/domain"
	"github.com/google/uuid"
)

func (s *Storage) GetRefreshToken(userId uuid.UUID) (domain.RefreshToken, error) {
	var refreshToken domain.RefreshToken
	result := s.db.Where("user_id = ?", userId).First(&refreshToken)
	if result.Error != nil {
		return refreshToken, result.Error
	}
	return refreshToken, nil
}

func (s *Storage) UpdateRefreshToken(token domain.RefreshToken) error {
	existingToken := domain.RefreshToken{}
	result := s.db.Where("user_id = ?", token.User.ID).First(&existingToken)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			result = s.db.Create(&token)
			if result.Error != nil {
				return result.Error
			} else {
				return nil
			}
		}
		return result.Error
	}
	result = s.db.Model(&existingToken).Updates(token)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
