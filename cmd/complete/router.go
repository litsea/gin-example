package complete

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	api "github.com/litsea/gin-api"
	"github.com/litsea/gin-api/errcode"
	i18n "github.com/litsea/gin-i18n"
	log "github.com/litsea/log-slog"
	"github.com/spf13/viper"

	"github.com/litsea/gin-example/config"
)

var (
	errTest     = errcode.New(1001, "errTest")
	errNotAdmin = errcode.New(1002, "errNotAdmin")
)

func newRouter(r *gin.Engine, v *viper.Viper) {
	r.GET("/", func(ctx *gin.Context) {
		api.Success(ctx, "OK")
	})

	r.GET("/user", func(ctx *gin.Context) {
		req := &struct {
			Name string `binding:"required,lte=10" form:"name"`
		}{}
		if err := ctx.ShouldBind(req); err != nil {
			api.VError(ctx, err, req)
			return
		}

		api.Success(ctx, req.Name)
	})

	r.GET("/panic", func(ctx *gin.Context) {
		panic("unknown panic")
	})

	r.GET("/no-translate", func(ctx *gin.Context) {
		api.Success(ctx, i18n.T(ctx, "NoTranslate"))
	})

	r.GET("/err-test", func(ctx *gin.Context) {
		log.Info("err-test", "req", ctx.Request)

		api.Error(ctx, errTest)
	})

	r.GET("/err-unknown", func(ctx *gin.Context) {
		log.Info("err-unknown", "req", ctx.Request)

		api.Error(ctx, fmt.Errorf("err-unknown"))
	})

	// Do not run on Windows git bash (GBK/UTF-8 issue)
	// curl -X POST -H 'Content-Type: application/json; charset=utf-8' \
	//   -d '{"name": "啊啊啊啊啊1111啊啊啊"}' http://localhost:8080/check-admin
	r.POST("/check-admin", func(ctx *gin.Context) {
		req := &struct {
			Name string `binding:"required,lte=20" json:"name"`
		}{}
		if err := ctx.ShouldBindJSON(req); err != nil {
			api.VError(ctx, err, req)
			return
		}

		if req.Name != "admin" {
			api.Error(ctx, errNotAdmin)
		} else {
			api.Success(ctx, req.Name)
		}
	})

	r.GET("/long-time", func(ctx *gin.Context) {
		time.Sleep(10 * time.Second)
		api.Success(ctx, "long time ago")
	})

	r.GET("/rate-limit", IpLimiter.Middleware(), func(ctx *gin.Context) {
		api.Success(ctx, "rate-limit")
	})

	r.HandleMethodNotAllowed = true
	r.NoMethod(api.HandleMethodNotAllowed())
	r.NoRoute(api.HandleNotFound())
	api.RouteRegisterPprof(r, v.GetString(config.KeyPprofToken))
	r.GET("/v1/health", api.HandleHealthCheck())
}
