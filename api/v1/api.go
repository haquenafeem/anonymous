package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/haquenafeem/anonymous/service"
)

type Api struct {
	router *gin.Engine
	svc    *service.Service
}

func (api *Api) setupUserRoutes(v1 *gin.RouterGroup) {
	users := v1.Group("/users")
	users.POST("/register", api.register)
	users.POST("/login", api.login)
	users.POST("/upload", api.Authenticate(), api.upload)
	users.GET("/generate-qr-code", api.Authenticate(), api.generateQRCode)
}

func (api *Api) setupMessageRoutes(v1 *gin.RouterGroup) {
	messages := v1.Group("/messages")
	messages.POST("/", api.postMessage)
	messages.GET("/", api.Authenticate(), api.getMessages)
	messages.GET("/share/:msg_id", api.Authenticate(), api.shareMessage)
}

func (api *Api) setupWebPages(engine *gin.Engine) {
	engine.GET("/", api.indexPage)
	engine.GET("/login", api.loginPage)
	engine.GET("/register", api.registerPage)
	engine.GET("/messages/:user_id", api.postMessagePage)
	engine.GET("/dashboard", api.dashboardPage)
	engine.GET("/404", api.notFound404)
}

func (api *Api) setupRoutes(v1 *gin.RouterGroup) {
	api.setupUserRoutes(v1)
	api.setupMessageRoutes(v1)
}

func (api *Api) Run(port int) error {
	if err := api.router.Run(fmt.Sprintf(":%d", port)); err != nil {
		return err
	}

	return nil
}

func New(engine *gin.Engine, svc *service.Service) *Api {
	api := &Api{
		router: engine,
		svc:    svc,
	}

	api.router.Static("/assets", "./assets")
	api.router.Static("/img", "./images")
	api.router.LoadHTMLGlob("templates/*")

	api.setupWebPages(api.router)
	v1 := api.router.Group("/api/v1")
	api.setupRoutes(v1)

	return api
}
