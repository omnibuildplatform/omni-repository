package common

import (
	"context"
	"github.com/omnibuildplatform/omni-repository/common/config"
	"github.com/omnibuildplatform/omni-repository/common/database"
	"github.com/omnibuildplatform/omni-repository/common/models"
	"github.com/omnibuildplatform/omni-repository/common/storage"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Store struct {
	Config   *config.PersistentStore
	Logger   *zap.Logger
	Database *gorm.DB
}

func NewStore(config *config.PersistentStore, logger *zap.Logger) (*Store, error) {
	database, err := database.ConnectToDB(config.User, config.Password, config.Host, config.Port, config.DBName)
	if err != nil {
		logger.Error("failed to initialize database object")
		return nil, err
	}
	db, err := database.DB()
	if err != nil {
		logger.Error("failed to get to database instance")
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		logger.Error("failed to connect to database")
		return nil, err
	}

	err = database.AutoMigrate(models.Image{})
	if err != nil {
		logger.Error("failed to auto migrate image model")
		return nil, err
	}
	return &Store{
		Config:   config,
		Logger:   logger,
		Database: database,
	}, nil
}

func (s *Store) GetImageStorage(ctx context.Context) *storage.ImageStorage {
	return storage.NewImageStorage(s.Database, ctx)
}

func (s *Store) Close() {
}
