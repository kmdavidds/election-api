package mysql

import (
	"log"

	"github.com/kmdavidds/election-api/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&entity.User{},
	)

	if err != nil {
		log.Fatalf("failed migration db: %v", err)
	}
}
