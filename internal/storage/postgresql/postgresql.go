package postgresql

import (
	"PartyRoom.API/internal/domain"
	"PartyRoom.API/internal/lib/constants/constants"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type Storage struct {
	db *gorm.DB
}

func New(connectionString string) (*Storage, error) {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database! \n", err.Error())
		os.Exit(1)
	}

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations")
	err = db.AutoMigrate(&domain.User{})
	err = db.AutoMigrate(&domain.Role{})
	err = db.AutoMigrate(&domain.UserRole{})
	err = db.AutoMigrate(&domain.RefreshToken{})
	err = db.AutoMigrate(&domain.Tag{})
	if err != nil {
		log.Fatal("Migration Failed: \n", err.Error())
		os.Exit(1)
	}
	// Создание трех ролей при первом запуске
	createDefaultRoles(db)
	log.Println("🚀 Connected Successfully to the Database")

	return &Storage{db: db}, nil
}
func createDefaultRoles(db *gorm.DB) {
	roles := []domain.Role{
		{Name: constants.UserRole.Name, UpperName: constants.UserRole.UpperName, LowerName: constants.UserRole.LowerName},
		{Name: constants.AdminRole.Name, UpperName: constants.AdminRole.UpperName, LowerName: constants.AdminRole.LowerName},
		{Name: constants.ModeratorRole.Name, UpperName: constants.ModeratorRole.UpperName, LowerName: constants.ModeratorRole.LowerName},
	}

	for _, role := range roles {
		var existingRole domain.Role
		result := db.Where("name = ?", role.Name).First(&existingRole)
		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
			log.Fatal("Error checking existing role: ", result.Error)
			os.Exit(1)
		}

		// Создавать роль, только если она еще не существует
		if result.Error == gorm.ErrRecordNotFound {
			result := db.Create(&role)
			if result.Error != nil {
				log.Fatal("Error creating default role: ", result.Error)
				os.Exit(1)
			}
		}
	}
}
