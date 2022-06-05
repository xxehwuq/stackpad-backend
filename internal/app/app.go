package app

import (
	"fmt"
	"log"

	"github.com/yaroslavyarosh/stackpad-backend/config"
	"github.com/yaroslavyarosh/stackpad-backend/internal/entity"
	"github.com/yaroslavyarosh/stackpad-backend/internal/service"
	"github.com/yaroslavyarosh/stackpad-backend/internal/storage"
	"github.com/yaroslavyarosh/stackpad-backend/internal/transport"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Run(cfg *config.Config) {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", cfg.Db.Username, cfg.Db.Password, cfg.Db.Host, cfg.Db.Port, cfg.Db.Name)), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}

	db.AutoMigrate(&entity.User{})

	storage := storage.New(db)
	service := service.New(storage)
	transport := transport.New(service)

	transport.Init(cfg)
}
