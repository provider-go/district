package api

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/district/models"
	"github.com/provider-go/pkg/output"
)

func ListDistrict(ctx *gin.Context) {
	upId := output.ParamToInt(ctx.Query("upId"))
	list, err := models.ListDistrict(upId)

	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, list)
	}
}
