package service

import (
	"github.com/haquenafeem/anonymous/internal"
	"github.com/haquenafeem/anonymous/repository"
)

type Service struct {
	repo *repository.Repository
	sm   *internal.ShareMessage
}

func New(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
		sm:   internal.NewShareMessage(),
	}
}
