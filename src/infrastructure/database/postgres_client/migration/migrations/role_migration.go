package migrations

import (
	"github.com/krostyle/auth-systme-go/src/infrastructure/database/postgres_client/model"
	"gorm.io/gorm"
)

func MigrateRoles(db *gorm.DB) error {
	return db.AutoMigrate(&model.Role{})
}
