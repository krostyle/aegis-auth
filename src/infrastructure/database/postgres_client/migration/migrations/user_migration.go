package migrations

import (
	"github.com/krostyle/auth-systme-go/src/infrastructure/database/postgres_client/model"
	"gorm.io/gorm"
)

func MigrateUsers(db *gorm.DB) error {
	return db.AutoMigrate(&model.User{})
}
