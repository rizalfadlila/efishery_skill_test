package boostrap

import (
	"github.com/fetch_app/constants"
	"github.com/fetch_app/docs"
	rest "github.com/fetch_app/handlers"
	"github.com/fetch_app/handlers/middlewares"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func initREST() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	pprof.Register(router)

	// Swagger
	docs.SwaggerInfo.Title = constants.ServiceName
	docs.SwaggerInfo.Version = constants.ServiceVersion

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(middlewares.CORSMiddleware())

	handler := rest.NewHandler(svc)

	authRouter := router.Use(middlewares.AuthMiddleware(), middlewares.UserContextMiddleware())
	authRouter.GET("/fetch", handler.Fetch)
	authRouter.GET("/clamis-jwt", handler.ClaimsJWT)

	adminRouter := router.Use(middlewares.AuthMiddleware(), middlewares.UserContextMiddleware(), middlewares.AdminMiddleware())
	adminRouter.GET("/aggregate", handler.Aggregate)

	return router
}
