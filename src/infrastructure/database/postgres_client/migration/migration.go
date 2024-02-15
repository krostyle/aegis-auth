package migration

import (
	"github.com/krostyle/auth-systme-go/src/infrastructure/database/postgres_client/migration/migrations"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) error {
	if err := migrations.MigratePermissions(db); err != nil {
		return err
	}
	return nil
}
