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
