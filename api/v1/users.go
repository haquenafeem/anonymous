package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haquenafeem/anonymous/model"
)

func (api *Api) register(ctx *gin.Context) {
	var req model.RegisterRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusExpectationFailed, &model.RegisterResponse{
			Err: "error binding json",
		})

		return
	}

	res := api.svc.RegisterUser(&req)
	if !res.IsSuccess {
		ctx.JSON(http.StatusInternalServerError, res)

		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (api *Api) login(ctx *gin.Context) {
	var req model.LoginRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusExpectationFailed, &model.LoginResponse{
			Err: "error binding json",
		})

		return
	}

	res := api.svc.LoginUser(&req)
	if !res.IsSuccess {
		ctx.JSON(http.StatusInternalServerError, res)

		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (api *Api) upload(ctx *gin.Context) {
	ctx.JSON(api.svc.UploadProfilePic(ctx))
}

func (api *Api) generateQRCode(ctx *gin.Context) {
	imageBytes, err := api.svc.GenerateQrCodeBytes(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": "cannot generate qr code",
		})

		return
	}

	// Set the appropriate headers for the response
	ctx.Header("Content-Disposition", "attachment; filename=qrcode.png")
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Transfer-Encoding", "binary")

	ctx.Data(http.StatusOK, "application/octet-stream", imageBytes)
}

func (api *Api) getProfilePic(ctx *gin.Context) {
	ctx.JSON(api.svc.ProfilePic(ctx))
}
