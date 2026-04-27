package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rheadavin/hr-go-api/internal/config"
	"github.com/rheadavin/hr-go-api/internal/handler"
	"github.com/rheadavin/hr-go-api/internal/middleware"
	"github.com/rheadavin/hr-go-api/internal/repository"
	"github.com/rheadavin/hr-go-api/internal/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.CORS())
	r.Use(middleware.CustomLogger())

	if os.Getenv("APP_ENV") != "production" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// dependency injection
	// repositories
	userRepo := repository.NewUserRepository(db)
	divisionRepo := repository.NewDivisionRepository(db)
	employeeRepo := repository.NewEmployeeRepository(db)

	// services
	authServices := service.NewAuthService(userRepo)
	divisionService := service.NewDivisionService(divisionRepo)
	employeeService := service.NewEmployeeService(employeeRepo)

	// handlers
	authHandler := handler.NewAuthHandler(authServices)
	divisionHandler := handler.NewDivisionHandler(divisionService)
	employeeHandler := handler.NewEmployeeHandler(employeeService)

	// routes
	api := r.Group("/api")

	// Health check
	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": config.AppConfig.AppName})
	})

	// Auth routes (public)
	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	protected := api.Group("")
	protected.Use(middleware.Auth())
	{
		protected.GET("/me", authHandler.Me)
	}

	// division group
	division := api.Group("/division")
	division.Use(middleware.Auth())
	{
		division.POST("/", divisionHandler.FindAll)
		division.GET("/:id", divisionHandler.FindByID)
		division.POST("/create", divisionHandler.Create)
		division.PUT("/:id", divisionHandler.Update)
		division.DELETE("/:id", divisionHandler.Delete)
	}

	// employee group
	employee := api.Group("/employee")
	employee.Use(middleware.Auth())
	{
		employee.POST("/", employeeHandler.FindAll)
		employee.POST("/create", employeeHandler.Create)
		employee.GET("/:id", employeeHandler.FindByID)
		employee.PUT("/:id", employeeHandler.Update)
		employee.DELETE("/:id", employeeHandler.Delete)
	}

	return r
}
