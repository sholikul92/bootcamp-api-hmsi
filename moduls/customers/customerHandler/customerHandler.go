package customerHandler

import (
	"bootcamp-api-hmsi/moduls/customers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	UC customers.CustomersUseCase
}

func NewCustomerHandler(r *gin.Engine, UC customers.CustomersUseCase) {
	handler := customerHandler{UC}

	r.GET("/customers", handler.GetAll)
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
