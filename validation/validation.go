package validation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CheckErrorType(c *gin.Context, err error) {
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		var messages []string
		for _, fieldError := range validationErrs {
			messages = append(messages, fieldError.Error()) // Thông báo lỗi từ tag validator
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": messages})
		return
	}
}
