package repository

import "github.com/haquenafeem/anonymous/model"

func (repository *Repository) CreateMessage(message *model.Message) error {
	return repository.db.Create(message).Error
}

func (repository *Repository) GetAll(userID string) ([]model.Message, error) {
	var messages []model.Message
	err := repository.db.Where("user_id = ?", userID).Find(&messages).Error

	return messages, err
}

func (repository *Repository) GetMessage(id string) (*model.Message, error) {
	var message model.Message
	err := repository.db.Where("id = ?", id).Find(&message).Error

	return &message, err
}
