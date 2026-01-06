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

func (h *Handler) GetSalaryEstimationWithReviews(c *gin.Context) {
	ctx := c.Request.Context()

	results, err := h.service.GetSalaryEstimationWithReviews(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, standard.Response{
			Success: false,
			Message: "Failed to calculate salary estimation",
		})
		return
	}

	c.JSON(http.StatusOK, standard.Response{
		Success: true,
		Message: "Success to calculate salary estimation",
		Result:  results,
	})
}

func (h *Handler) SaveActiveSmithEmployeesToFile(c *gin.Context) {
	ctx := c.Request.Context()

	data, err := h.service.GetActiveSmithEmployees(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, standard.Response{
			Success: false,
			Message: "Failed to fetch data employees",
		})
		return
	}

	filename := "contoh2.txt"

	err = SaveJSONToFile(filename, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, standard.Response{
			Success: false,
			Message: "Failed to save file",
		})
		return
	}

	c.JSON(http.StatusOK, standard.Response{
		Success: true,
		Message: "File save successfully",
	})
}
func (h *Handler) SaveEmployeesWithoutReviewsToFile(c *gin.Context) {
	ctx := c.Request.Context()

	data, err := h.service.GetEmployeesWithoutReviews(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, standard.Response{
			Success: false,
			Message: "Failed to fetch data employees",
		})
		return
	}

	filename := "contoh3.txt"

	err = SaveJSONToFile(filename, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, standard.Response{
			Success: false,
			Message: "Failed to save file",
		})
		return
	}

	c.JSON(http.StatusOK, standard.Response{
		Success: true,
		Message: "File save successfully",
	})
}

func (h *Handler) SaveHireDateDiffActiveEmployeesToFile(c *gin.Context) {
	ctx := c.Request.Context()

	data, err := h.service.GetHireDateDiffActiveEmployees(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, standard.Response{
			Success: false,
			Message: "Failed to calculate hire date difference",
		})
		return
	}

	filename := "contoh4.txt"

	err = SaveJSONToFile(filename, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, standard.Response{
			Success: false,
			Message: "Failed to save file",
		})
		return
	}

	c.JSON(http.StatusOK, standard.Response{
		Success: true,
		Message: "File save successfully",
	})
}

func (h *Handler) SaveSalaryEstimationWithReviewsToFile(c *gin.Context) {
	ctx := c.Request.Context()

	data, err := h.service.GetSalaryEstimationWithReviews(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, standard.Response{
			Success: false,
			Message: "Failed to calculate salary estimation",
		})
		return
	}

	filename := "contoh5.txt"

	err = SaveJSONToFile(filename, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, standard.Response{
			Success: false,
			Message: "Failed to save file",
		})
		return
	}

	c.JSON(http.StatusOK, standard.Response{
		Success: true,
		Message: "File save successfully",
	})
}
