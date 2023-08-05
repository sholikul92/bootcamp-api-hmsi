package customerHandler

import (
	"bootcamp-api-hmsi/models"
	"bootcamp-api-hmsi/moduls/customers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	UC customers.CustomersUseCase
}

func NewCustomerHandler(r *gin.Engine, UC customers.CustomersUseCase) {
	handler := customerHandler{UC}

	r.GET("/customers", handler.GetAll)
	r.POST("/customers", handler.Insert)
	r.GET("/customers/:id", handler.FindById)
	r.PUT("/customers/:id", handler.UpdateById)

}

func (h *customerHandler) GetAll(c *gin.Context) {
	result, err := h.UC.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    []string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    result,
	})
}

func (h *customerHandler) Insert(c *gin.Context) {
	var request models.RequestInsertCustomer

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    []string{},
		})
		return
	}

	err := h.UC.Insert(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    []string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Inserted successfully",
		"data":    []string{},
	})
}

func (h *customerHandler) FindById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	customer, err := h.UC.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    customer,
	})
}

func (h *customerHandler) UpdateById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var customer models.RequestInsertCustomer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.UC.UpdateById(uint(id), &customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "updated successfuly",
	})
}
