package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pighead4u/options/src/model"
	"net/http"
)

/**
 * 根据交易所查询要显示的标题
 **/
func GetContentsByContractAndDate(context *gin.Context) {
	contract := context.Param("contract")
	date := context.Query("date")
	constractType := context.Query("type")

	fmt.Println(constractType)

	data := model.QueryContentByContractAndDateAndType(contract, date, constractType)

	var response model.OptionResponse
	response.Data = data
	response.IsSuccess = true

	context.JSON(http.StatusOK, response)
}
