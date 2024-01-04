package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haquenafeem/anonymous/model"
)

func (api *Api) postMessage(ctx *gin.Context) {
	var req model.PostMessageRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusExpectationFailed, &model.PostMessageResponse{
			Err: "error binding json",
		})

		return
	}

	res := api.svc.PostMessage(&req)
	if !res.IsSuccess {
		ctx.JSON(http.StatusInternalServerError, res)

		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (api *Api) getMessages(ctx *gin.Context) {
	ctx.JSON(api.svc.GetAllMessages(ctx))
}

func (api *Api) shareMessage(ctx *gin.Context) {

}
