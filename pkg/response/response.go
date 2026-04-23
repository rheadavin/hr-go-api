package response

import "github.com/gin-gonic/gin"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
	Meta    *MetaData   `json:"meta,omitempty"`
}
type MetaData struct {
	Total int64 `json:"total"`
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Pages int   `json:"pages"`
}

func SuccessResponse(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(code, Response{
		Success: true,
		Message: msg,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, msg string) {
	c.JSON(code, Response{
		Success: false,
		Message: msg,
	})
}

func PaginatedResponse(c *gin.Context, data interface{}, meta *MetaData) {
	c.JSON(200, Response{
		Success: true,
		Message: "OK",
		Data:    data,
		Meta:    meta,
	})
}
