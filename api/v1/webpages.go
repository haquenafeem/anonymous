package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *Api) indexPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func (api *Api) loginPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
}

func (api *Api) registerPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register.html", nil)
}

func (api *Api) postMessagePage(ctx *gin.Context) {

}
