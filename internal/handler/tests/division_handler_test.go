package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rheadavin/hr-go-api/internal/dto"
	"github.com/rheadavin/hr-go-api/internal/handler"
	"github.com/rheadavin/hr-go-api/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupRouter(h *handler.DivisionHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/division", h.FindAll)
	r.POST("/division/create", h.Create)
	r.GET("/division/:id", h.FindByID)
	r.PUT("/division/:id", h.Update)
	r.DELETE("/division/:id", h.Delete)
	return r
}

func TestDivisionHandler_FindAll(t *testing.T) {
	mockSvc := new(mocks.DivisionServiceInterface)

	mockSvc.On("FindAll", 1, 10, "").Return([]dto.DivisionResponse{
		{ID: 1, Name: "Engineering"},
		{ID: 2, Name: "HR"},
	}, int64(2), nil)

	h := handler.NewDivisionHandler(mockSvc)
	r := setupRouter(h)
	// Buat request
	reqBody := bytes.NewBufferString(`{"page": 1, "limit": 10, "search": ""}`)
	req, _ := http.NewRequest(http.MethodPost, "/division", reqBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// assert response
	assert.Equal(t, http.StatusOK, w.Code)
	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)
	assert.True(t, body["success"].(bool))
	mockSvc.AssertExpectations(t)
}

func TestDivisionHandler_Create_ValidationError(t *testing.T) {
	mockSvc := new(mocks.DivisionServiceInterface)
	mockSvc.On("Create", mock.Anything).Return(nil, nil).Maybe()

	h := handler.NewDivisionHandler(mockSvc)
	r := setupRouter(h)
	// Kirim body yang tidak valid (name kosong)
	body := bytes.NewBufferString(`{"name": "", "description": ""}`)

	req, _ := http.NewRequest(http.MethodPost, "/division/create", body)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	// Service tidak boleh dipanggil kalau validasi gagal
	mockSvc.AssertNotCalled(t, "Create")
}
