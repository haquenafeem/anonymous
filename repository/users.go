package repository

import (
	"github.com/haquenafeem/anonymous/model"
)

func (repository *Repository) CreateUser(user *model.User) error {
	return repository.db.Create(user).Error
}

func (repository *Repository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := repository.db.Where("email = ?", email).Find(&user).Error

	return &user, err
}

func (repository *Repository) FindUser(id string) (*model.User, error) {
	var user model.User
	err := repository.db.Where("id = ?", id).Find(&user).Error

	return &user, err
}

func (repository *Repository) UpdateProfilePicId(userId, picId string) *model.UploadResponse {
	query := `
	UPDATE
		users
	SET
		profile_pic_id = ?
	Where
		id = ?
	`

	err := repository.db.Exec(query, picId, userId).Error
	if err != nil {
		return &model.UploadResponse{
			Err: "db update failed",
		}
	}

	return &model.UploadResponse{
		IsSuccess: true,
	}
}
