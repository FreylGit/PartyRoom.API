package postgresql

import (
	"PartyRoom.API/internal/domain"
	"github.com/google/uuid"
)

func (s *Storage) SaveTag(tag domain.Tag) error {
	result := s.db.Create(&tag)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Storage) DeleteTag(tagID uuid.UUID, userID uuid.UUID) error {
	result := s.db.Delete(&domain.Tag{}, "id = ? and user_id = ?", tagID, userID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Storage) UpdateTag(tag domain.Tag) error {
	existingTag := domain.Tag{}
	result := s.db.Where("id = ? and user_id = ?", tag.ID, tag.UserID).First(&existingTag)
	if result.Error != nil {
		return result.Error
	}
	result = s.db.Model(&existingTag).Updates(tag)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
