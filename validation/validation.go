package validation

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nhutHao02/social-network-common-service/model"
	constants "github.com/nhutHao02/social-network-common-service/utils/constanst"
)

func CheckErrorType(c *gin.Context, err error) {
	switch err.(type) {
	case *json.UnmarshalTypeError:
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(err, constants.RequestInvalid))
		return
	case *json.SyntaxError:
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(err, constants.RequestInvalid))
		return
	case validator.ValidationErrors:
		HandleValidationErrors(c, err)
		return
	default:
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(err, constants.RequestInvalid))
		return
	}

}
func HandleValidationErrors(c *gin.Context, err error) {
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		var messages []string
		for _, fieldError := range validationErrs {
			messages = append(messages, fieldError.Error())
		}
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(err, strings.Join(messages, ",")))
		return
	}
}
