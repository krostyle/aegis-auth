package migrations

import (
	"github.com/krostyle/auth-systme-go/src/infrastructure/database/postgres_client/model"
	"gorm.io/gorm"
)

func MigratePermissions(db *gorm.DB) error {
	return db.AutoMigrate(&model.Permission{})
}
