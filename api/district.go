package api

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/district/models"
	"github.com/provider-go/pkg/output"
)

func ListDistrict(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	upId := output.ParamToInt(json["upId"])
	list, err := models.ListDistrict(upId)

	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, list)
	}
}
