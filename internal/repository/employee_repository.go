package repository

import (
	"github.com/rheadavin/hr-go-api/internal/dto"
	"github.com/rheadavin/hr-go-api/internal/models"
	"gorm.io/gorm"
)

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
	var employees []dto.EmployeeResponse
	var total int64
	query := r.db.Model(&models.Employee{}).Select("id", "nik", "full_name", "email", "is_active", "division_id").Preload("Division").Where("is_active = ?", true)

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
		Scan(&employees).Error

	return employees, total, err
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
