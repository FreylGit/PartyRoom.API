package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

const (
	RoleNameUser      = "User"
	RoleNameAdmin     = "Admin"
	RoleNameModerator = "Moderator"
)

type User struct {
	ID               *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key;"`
	Email            string     `gorm:"type:varchar(255);not null;unique"`
	PasswordHash     string     `gorm:"type:varchar(255);not null"`
	Name             string     `gorm:"type:varchar(255);not null"`
	Phone            string     `gorm:"type:varchar(100)"`
	Photo            string     `gorm:"type:varchar(255)"`
	IsConfirmedEmail bool       `gorm:"type:boolean;default:false"`
	Roles            []*Role    `gorm:"many2many:user_roles;"`
	RefreshToken     RefreshToken
}

type Role struct {
	ID        *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key;"`
	Name      string     `gorm:"type:varchar(255);not null;unique"`
	UpperName string     `gorm:"type:varchar(255)"`
	LowerName string     `gorm:"type:varchar(255)"`
}

type UserRole struct {
	gorm.Model
	UserID *uuid.UUID `gorm:"type:uuid;index"`
	RoleID *uuid.UUID `gorm:"type:uuid;index"`
	User   *User      `gorm:"foreignKey:UserID"`
	Role   *Role      `gorm:"foreignKey:RoleID"`
}

type RefreshToken struct {
	ID             *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key;"`
	UserID         *uuid.UUID `gorm:"type:uuid;index"`
	User           *User      `gorm:"foreignKey:UserID"`
	Token          string     `gorm:"type:varchar(255);not null"`
	ExpirationDate time.Time  `gorm:"type:time"`
	CreateDate     string     `grom:"type:time"`
}

func (token *RefreshToken) BeforeCreate(tx *gorm.DB) (err error) {
	// Устанавливаем ExpirationDate на текущую дату плюс три дня
	token.ExpirationDate = time.Now().Add(3 * 24 * time.Hour)
	return nil
}
