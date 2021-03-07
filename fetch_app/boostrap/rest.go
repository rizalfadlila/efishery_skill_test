package boostrap

import (
	"os"

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
	if os.Getenv("ENV") != "production" {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	router.Use(middlewares.CORSMiddleware())

	handler := rest.NewHandler(svc)

	router.GET("/fetch", handler.Fetch)
	router.GET("/aggregate", handler.Aggregate)
	router.GET("/clamis-jwt", handler.ClaimsJWT)

	return router
}
