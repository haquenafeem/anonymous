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
	imageBytes, err := api.svc.ShareMessage(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": "cannot generate image for message",
		})

		return
	}

	// Set the appropriate headers for the response
	ctx.Header("Content-Disposition", "attachment; filename=message.png")
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Transfer-Encoding", "binary")

	ctx.Data(http.StatusOK, "application/octet-stream", imageBytes)
}
