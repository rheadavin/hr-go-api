package service

import (
	"errors"

	"github.com/rheadavin/hr-go-api/internal/dto"
	"github.com/rheadavin/hr-go-api/internal/models"
	"github.com/rheadavin/hr-go-api/internal/repository"
)

type DivisionServiceInterface interface {
	Create(req dto.CreateDivisionRequest) (*dto.DivisionResponse, error)
	FindAll(page, limit int, search string) ([]dto.DivisionResponse, int64, error)
	FindByID(id uint) (*models.Division, error)
	Update(id uint, req dto.UpdateDivisionRequest) (*dto.DivisionResponse, error)
	Delete(id uint) error
}

type DivisionService struct {
	divisionRepo *repository.DivisionRepository
}

func NewDivisionService(divisionRepo *repository.DivisionRepository) *DivisionService {
	return &DivisionService{divisionRepo: divisionRepo}
}

func (s *DivisionService) Create(req dto.CreateDivisionRequest) (*dto.DivisionResponse, error) {
	division := models.Division{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := s.divisionRepo.Create(&division); err != nil {
		return nil, errors.New("failed to create division")
	}

	return &dto.DivisionResponse{
		ID:   division.ID,
		Name: division.Name,
	}, nil
}

func (s *DivisionService) FindAll(page, limit int, search string) ([]dto.DivisionResponse, int64, error) {
	divisions, total, err := s.divisionRepo.FindAll(page, limit, search)
	if err != nil {
		return nil, 0, errors.New("failed to find divisions")
	}

	return divisions, total, nil
}

func (s *DivisionService) FindByID(id uint) (*models.Division, error) {
	return s.divisionRepo.FindByID(id)
}

func (s *DivisionService) Update(id uint, req dto.UpdateDivisionRequest) (*dto.DivisionResponse, error) {
	division := models.Division{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := s.divisionRepo.Update(id, &division); err != nil {
		return nil, errors.New("failed to update division")
	}

	return &dto.DivisionResponse{
		ID:   division.ID,
		Name: division.Name,
	}, nil
}

func (s *DivisionService) Delete(id uint) error {
	return s.divisionRepo.Delete(id)
}
