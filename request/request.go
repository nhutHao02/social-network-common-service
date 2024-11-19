package request

import (
	"github.com/gin-gonic/gin"
	"github.com/nhutHao02/social-network-common-service/model"
	constants "github.com/nhutHao02/social-network-common-service/utils/constanst"
	"github.com/nhutHao02/social-network-common-service/utils/logger"
	"github.com/nhutHao02/social-network-common-service/validation"
)

func GetBodyJSON(c *gin.Context, dest interface{}) (err error) {
	err = c.ShouldBindJSON(dest)
	if err != nil {
		logger.Error(constants.ShoulBindJsonError)
		validation.CheckErrorType(c, err)
		return
	}
	return
}

func GetParam(c *gin.Context, key string) (data string) {
	data = c.Param(key)
	return
}

// add tag uri in dest struct
func GetParamsFromUrl(c *gin.Context, dest interface{}) (err error) {
	err = c.ShouldBindUri(dest)
	if err != nil {
		logger.Error(constants.ShoulBindJsonError)
		validation.CheckErrorType(c, err)
		return
	}
	return
}

// add tag form in dest struct
func GetQueryParamsFromUrl(c *gin.Context, dest interface{}) (err error) {
	err = c.ShouldBindQuery(dest)
	if err != nil {
		logger.Error(constants.ShoulBindJsonError)
		validation.CheckErrorType(c, err)
		return
	}
	return
}

func GetPaging(c *gin.Context) *model.Paging {
	var p model.Paging
	err := c.ShouldBindQuery(&p)
	if err != nil {
		logger.Error(constants.ShoulBindJsonError)
		validation.CheckErrorType(c, err)
		return nil
	}
	return &p
}
