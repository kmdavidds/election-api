package mysql

import (
	"log"

	"github.com/kmdavidds/election-api/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&entity.User{},
		&entity.Post{},
		&entity.Comment{},
	)

	if err != nil {
		log.Fatalf("failed migration db: %v", err)
	}
}
