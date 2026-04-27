package handler

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rheadavin/hr-go-api/internal/dto"
	"github.com/rheadavin/hr-go-api/internal/service"
	response "github.com/rheadavin/hr-go-api/pkg/response"
)

type EmployeeHandler struct {
	empService *service.EmployeeService
}

func NewEmployeeHandler(empService *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{empService: empService}
}

func (h *EmployeeHandler) FindAll(c *gin.Context) {
	var req dto.ListEmployeeRequest
	req.Page = 1
	req.Limit = 10
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	employees, total, err := h.empService.FindAll(req.Page, req.Limit, req.Search)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.PaginatedResponse(c, employees, &response.MetaData{
		Total: total,
		Page:  req.Page,
		Limit: req.Limit,
		Pages: int(math.Ceil(float64(total) / float64(req.Limit))),
	})
}

func (h *EmployeeHandler) Create(c *gin.Context) {
	var req dto.CreateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.empService.Create(req)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusCreated, "Employee created successfully", result)
}

func (h *EmployeeHandler) FindByID(c *gin.Context) {
	empId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid employee ID")
		return
	}

	employee, err := h.empService.FindByID(uint(empId))
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Employee retrieved successfully", employee)
}

func (h *EmployeeHandler) Update(c *gin.Context) {
	empId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid employee ID")
		return
	}

	var req dto.UpdateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.empService.Update(uint(empId), req)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Employee updated successfully", result)
}

func (h *EmployeeHandler) Delete(c *gin.Context) {
	empId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid employee ID")
		return
	}

	err = h.empService.Delete(uint(empId))
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Employee deleted successfully", nil)
}
