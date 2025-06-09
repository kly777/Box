package service

import (
	"box/internal/storage/repos"
	"box/internal/storage/models"
)

type BoxService interface {
	GetRootBoxes() ([]models.Box, error)
	CreateBox(name string) (*models.Box, error)
}

type LocalBoxService struct{}

func (s *LocalBoxService) GetRootBoxes() ([]models.Box, error) {
	return crud.GetRootBoxes()
}

func (s *LocalBoxService) CreateBox(name string) (*models.Box, error) {
	box := &models.Box{Name: name}
	err := crud.CreateBox(box)
	return box, err
}
