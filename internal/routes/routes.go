package routes

import (
	"be-employee-management/internal/database"
	"be-employee-management/internal/employee"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {
	employeeRepo := employee.NewRepository(database.DB)
	employeeService := employee.NewService(employeeRepo)
	employeeHandler := employee.NewHandler(employeeService)

	apiV1 := r.Group("/api/v1")
	apiV1.GET("/employees/active/smith", employeeHandler.GetActiveSmithEmployees)
	apiV1.GET("/employees/no-reviews", employeeHandler.GetEmployeesWithoutReviews)
	apiV1.GET("/employees/hire-date-diff", employeeHandler.GetHireDateDiffActiveEmployees)
	apiV1.GET("/employees/salary-estimation", employeeHandler.GetSalaryEstimationWithReviews)
}
