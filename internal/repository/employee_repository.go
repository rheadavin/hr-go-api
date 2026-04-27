package repository

import (
	"github.com/rheadavin/hr-go-api/internal/dto"
	"github.com/rheadavin/hr-go-api/internal/models"
	"gorm.io/gorm"
)

type EmployeeRepositoryInterface interface {
	Create(employee *models.Employee) error
	FindAll(page, limit int, search string) ([]dto.EmployeeResponse, int64, error)
	FindByID(id uint) (*models.Employee, error)
	Update(id uint, data *models.Employee) error
	Delete(id uint) error
}

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

// create
func (r *EmployeeRepository) Create(employee *models.Employee) error {
	if err := r.db.Create(employee).Error; err != nil {
		return err
	}

	return r.db.Preload("Division").First(employee, employee.ID).Error
}

// get all
func (r *EmployeeRepository) FindAll(page, limit int, search string) ([]dto.EmployeeResponse, int64, error) {
	var employees []models.Employee
	var total int64
	query := r.db.Model(&models.Employee{}).
		Preload("Division").
		Where("is_active = ?", true)

	// Filter search
	if search != "" {
		query = query.Where(
			"full_name ILIKE ? OR nik ILIKE ? OR email ILIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%",
		)
	}

	// Count total
	query.Count(&total)

	// Pagination
	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).
		Order("created_at DESC").
		Find(&employees).Error

	if err != nil {
		return nil, 0, err
	}

	var result []dto.EmployeeResponse
	for _, emp := range employees {
		result = append(result, dto.EmployeeResponse{
			ID:         emp.ID,
			NIK:        emp.NIK,
			FullName:   emp.FullName,
			Email:      emp.Email,
			DivisionID: emp.DivisionID,
			Division: dto.DivisionResponse{
				ID:   emp.Division.ID,
				Name: emp.Division.Name,
			},
		})
	}

	return result, total, nil
}

// get by id
func (r *EmployeeRepository) FindByID(id uint) (*models.Employee, error) {
	var employee models.Employee
	err := r.db.Preload("Division").First(&employee, id).Error
	if err != nil {
		return nil, err // gorm.ErrRecordNotFound jika tidak ada
	}
	return &employee, nil
}

// update
func (r *EmployeeRepository) Update(id uint, data *models.Employee) error {
	if err := r.db.Model(&models.Employee{}).Where("id = ? AND is_active = ?", id, true).Updates(data).Error; err != nil {
		return err
	}

	return r.db.Preload("Division").First(data, id).Error
}

// delete
func (r *EmployeeRepository) Delete(id uint) error {
	return r.db.Model(&models.Employee{}).Where("id = ? AND is_active = ?", id, true).Update("is_active", false).Error
}
