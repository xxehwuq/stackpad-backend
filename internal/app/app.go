package app

import (
	"log"

	"github.com/yaroslavyarosh/stackpad-backend/config"
	"github.com/yaroslavyarosh/stackpad-backend/internal/service"
	"github.com/yaroslavyarosh/stackpad-backend/internal/storage"
	"github.com/yaroslavyarosh/stackpad-backend/internal/transport"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Run(cfg *config.Config) {
	db, err := gorm.Open(postgres.Open("user=postgres password=kJQ*Wo8MnNJfab5mnfjk host=db.rjxuletivavhwcvffefy.supabase.co port=5432 dbname=postgres"), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}

	storage := storage.New(db)
	service := service.New(storage)
	transport := transport.New(service)

	transport.Init(cfg)
}
