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

// @Summary Get all divisions
// @Description Get all divisions
// @Tags Division
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.ListDivisionRequest true "List division request"
// @Success 200 {object} response.Response{data=[]dto.DivisionResponse}
// @Failure 400 {object} response.Response
// @Router /division [post]
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

// @Summary Get division by ID
// @Description Get division by ID
// @Tags Division
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Division ID"
// @Success 200 {object} response.Response{data=dto.DivisionResponse}
// @Failure 400 {object} response.Response
// @Router /division/{id} [get]
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

// @Summary Create division
// @Description Create division
// @Tags Division
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateDivisionRequest true "Create division request"
// @Success 200 {object} response.Response{data=dto.DivisionResponse}
// @Failure 400 {object} response.Response
// @Router /division/create [post]
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

// @Summary Update division
// @Description Update division
// @Tags Division
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Division ID"
// @Param request body dto.UpdateDivisionRequest true "Update division request"
// @Success 200 {object} response.Response{data=dto.DivisionResponse}
// @Failure 400 {object} response.Response
// @Router /division/{id} [put]
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

// @Summary Delete division
// @Description Delete division
// @Tags Division
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Division ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /division/{id} [delete]
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
