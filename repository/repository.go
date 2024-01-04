package repository

import (
	"github.com/haquenafeem/anonymous/model"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewMust(db *gorm.DB) *Repository {
	repo := &Repository{
		db: db,
	}

	repo.migrateAllMust()

	return repo
}

func (repo *Repository) migrateAllMust() {
	err := repo.db.AutoMigrate(&model.User{}, &model.Message{})
	if err != nil {
		panic(err)
	}
}
