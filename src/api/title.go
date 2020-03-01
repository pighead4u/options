package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pighead4u/options/src/model"
	"net/http"
)

/**
 * 根据交易所查询要显示的标题
 **/
func GetTitlesByBourse(context *gin.Context) {
	name := context.Param("bourse")

	data := model.QueryTitleByBourse(name)

	var response model.OptionResponse
	response.Data = data
	response.IsSuccess = true

	context.JSON(http.StatusOK, response)
}
