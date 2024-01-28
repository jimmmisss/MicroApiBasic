package server

import (
	"github.com/gin-gonic/gin"
	"microApiBasic/internal/handler"
	"microApiBasic/internal/middleware"
	"microApiBasic/pkg/helper/resp"
	"microApiBasic/pkg/log"
)

func NewServerHTTP(
	logger *log.Logger,
	userHandler handler.UserHandler,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(
		middleware.CORSMiddleware(),
	)
	r.GET("/", func(ctx *gin.Context) {
		resp.HandleSuccess(ctx, map[string]interface{}{
			"say": "Hi Nunu!",
		})
	})
	r.GET("/user", userHandler.GetUserById)

	return r
}
