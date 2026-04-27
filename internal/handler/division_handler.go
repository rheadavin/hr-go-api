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

type DivisionHandler struct {
	divisionService service.DivisionServiceInterface
}

func NewDivisionHandler(divisionService service.DivisionServiceInterface) *DivisionHandler {
	return &DivisionHandler{divisionService: divisionService}
}

func (h *DivisionHandler) FindAll(c *gin.Context) {
	var req dto.ListDivisionRequest

	req.Page = 1
	req.Limit = 10

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	divisions, total, err := h.divisionService.FindAll(req.Page, req.Limit, req.Search)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.PaginatedResponse(c, divisions, &response.MetaData{
		Total: total,
		Page:  req.Page,
		Limit: req.Limit,
		Pages: int(math.Ceil(float64(total) / float64(req.Limit))),
	})
}

func (h *DivisionHandler) FindByID(c *gin.Context) {
	divisionId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid division ID")
		return
	}

	division, err := h.divisionService.FindByID(uint(divisionId))
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Division retrieved successfully", division)
}

func (h *DivisionHandler) Create(c *gin.Context) {
	var req dto.CreateDivisionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.divisionService.Create(req)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusCreated, "Division created successfully", result)
}

func (h *DivisionHandler) Update(c *gin.Context) {
	divisionId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid division ID")
		return
	}

	var req dto.UpdateDivisionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.divisionService.Update(uint(divisionId), req)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Division updated successfully", result)
}

func (h *DivisionHandler) Delete(c *gin.Context) {
	divisionId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid division ID")
		return
	}

	err = h.divisionService.Delete(uint(divisionId))
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Division deleted successfully", nil)
}
