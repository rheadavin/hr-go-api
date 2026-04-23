package service

import (
	"errors"

	"github.com/rheadavin/hr-go-api/internal/dto"
	"github.com/rheadavin/hr-go-api/internal/models"
	"github.com/rheadavin/hr-go-api/internal/repository"
)

type EmployeeService struct {
	empRepo *repository.EmployeeRepository
}

func NewEmployeeService(empRepo *repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{empRepo: empRepo}
}

func (s *EmployeeService) FindAll(page, limit int, search string) ([]dto.EmployeeResponse, int64, error) {
	return s.empRepo.FindAll(page, limit, search)
}

func (s *EmployeeService) Create(req dto.CreateEmployeeRequest) (*dto.EmployeeResponse, error) {
	employee := models.Employee{
		NIK:        req.NIK,
		FullName:   req.FullName,
		Email:      req.Email,
		Phone:      req.Phone,
		Position:   req.Position,
		Salary:     req.Salary,
		JoinDate:   req.JoinDate,
		DivisionID: req.DivisionID,
	}

	if err := s.empRepo.Create(&employee); err != nil {
		return nil, errors.New("failed to create employee")
	}

	return &dto.EmployeeResponse{
		ID:         employee.ID,
		NIK:        employee.NIK,
		FullName:   employee.FullName,
		Email:      employee.Email,
		DivisionID: employee.DivisionID,
		Division: dto.DivisionResponse{
			ID:   employee.Division.ID,
			Name: employee.Division.Name,
		},
	}, nil
}

func (s *EmployeeService) FindByID(id uint) (*models.Employee, error) {
	return s.empRepo.FindByID(id)
}

func (s *EmployeeService) Update(id uint, req dto.UpdateEmployeeRequest) (*dto.EmployeeResponse, error) {
	employee := models.Employee{
		NIK:        req.NIK,
		FullName:   req.FullName,
		Email:      req.Email,
		Phone:      req.Phone,
		Position:   req.Position,
		Salary:     req.Salary,
		JoinDate:   req.JoinDate,
		DivisionID: req.DivisionID,
	}

	if err := s.empRepo.Update(id, &employee); err != nil {
		return nil, errors.New("failed to update employee")
	}

	return &dto.EmployeeResponse{
		ID:         employee.ID,
		NIK:        employee.NIK,
		FullName:   employee.FullName,
		Email:      employee.Email,
		DivisionID: employee.DivisionID,
		Division: dto.DivisionResponse{
			ID:   employee.Division.ID,
			Name: employee.Division.Name,
		},
	}, nil
}

func (s *EmployeeService) Delete(id uint) error {
	return s.empRepo.Delete(id)
}
