package complete

import (
	"fmt"

	"github.com/gin-gonic/gin"
	api "github.com/litsea/gin-api"
	"github.com/litsea/gin-api/errcode"
	log "github.com/litsea/log-slog"
)

var errTest = errcode.New(1001, "errTest")

func newRouter(r *gin.Engine) {
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

	r.GET("/err-test", func(ctx *gin.Context) {
		log.Info("err-test", "req", ctx.Request)

		api.Error(ctx, errTest)
	})

	r.GET("/err-unknown", func(ctx *gin.Context) {
		log.Info("err-unknown", "req", ctx.Request)

		api.Error(ctx, fmt.Errorf("err-unknown"))
	})
}
