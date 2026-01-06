package employee

import (
	standard "be-employee-management/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetActiveSmithEmployees(c *gin.Context) {
	ctx := c.Request.Context()

	employees, err := h.service.GetActiveSmithEmployees(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, standard.Response{
			Success: false,
			Message: "Failed to fetch employees",
		})
		return
	}

	c.JSON(http.StatusOK, standard.Response{
		Success: true,
		Message: "Success to get employess",
		Result:  employees,
	})
}

func (h *Handler) GetEmployeesWithoutReviews(c *gin.Context) {
	ctx := c.Request.Context()

	employees, err := h.service.GetEmployeesWithoutReviews(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, standard.Response{
			Success: false,
			Message: "Failed to fetch employess",
		})
		return
	}

	c.JSON(http.StatusOK, standard.Response{
		Success: true,
		Message: "Success to get employess",
		Result:  employees,
	})
}

func (h *Handler) GetHireDateDiffActiveEmployees(c *gin.Context) {
	ctx := c.Request.Context()

	diffDays, err := h.service.GetHireDateDiffActiveEmployees(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, standard.Response{
			Success: false,
			Message: "Failed to calculate hire date difference",
		})
		return
	}

	c.JSON(http.StatusOK, standard.Response{
		Success: true,
		Message: "Success to calculate hire date difference",
		Result:  diffDays,
	})
}
