package employee

import (
	standard "be-employee-management/pkg/response"
	"encoding/json"
	"net/http"
	"path/filepath"

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
		Message: "Success to get employees",
		Result:  employees,
	})
}

func (h *Handler) GetEmployeesWithoutReviews(c *gin.Context) {
	ctx := c.Request.Context()

	employees, err := h.service.GetEmployeesWithoutReviews(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, standard.Response{
			Success: false,
			Message: "Failed to fetch employees",
		})
		return
	}

	c.JSON(http.StatusOK, standard.Response{
		Success: true,
		Message: "Success to get employees",
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

	filePath := filepath.Join("storage/reports", filename)

	c.Header("Content-Type", "text/plain")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.File(filePath)
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

	filePath := filepath.Join("storage/reports", filename)

	c.Header("Content-Type", "text/plain")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.File(filePath)
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

	filePath := filepath.Join("storage/reports", filename)

	c.Header("Content-Type", "text/plain")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.File(filePath)
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

	filePath := filepath.Join("storage/reports", filename)

	c.Header("Content-Type", "text/plain")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.File(filePath)
}

func (h *Handler) GetReportFromFile(c *gin.Context) {
	filename := c.Query("filename")
	if filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "filename is required",
		})
		return
	}

	bytes, err := ReadJSONFromFile(filename)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "file not found",
		})
		return
	}

	var jsonData any
	err = json.Unmarshal(bytes, &jsonData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, standard.Response{
			Success: false,
			Message: "Invalid json content",
		})
		return
	}

	c.JSON(http.StatusOK, standard.Response{
		Success: true,
		Message: "Success read file txt",
		Result:  jsonData,
	})
}
