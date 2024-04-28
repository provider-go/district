package api

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/district/models"
	"github.com/provider-go/pkg/ioput"
)

func ListDistrict(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	upId := ioput.ParamToInt(json["upId"])
	list, err := models.ListDistrict(upId)

	if err != nil {
		ioput.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		ioput.ReturnSuccessResponse(ctx, list)
	}
}
