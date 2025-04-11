package complete

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/litsea/gin-api/graceful"
	apilog "github.com/litsea/gin-api/log"
	log "github.com/litsea/log-slog"
	"github.com/spf13/viper"

	"github.com/litsea/gin-example/config"
)

func newServer(v *viper.Viper) {
	r := gin.New()

	// logger for gin-api, gin-i18n and other components/middlewares.
	l := apilog.New(
		log.GetSlog(),
		apilog.WithRequestHeader(true),
		apilog.WithRequestBody(true),
		apilog.WithUserAgent(true),
		apilog.WithStackTrace(true),
	)

	addMiddleware(r, v, l)
	newRouter(r, v)
	gracefulRunServer(v, r, l)
}

func gracefulRunServer(v *viper.Viper, r *gin.Engine, l apilog.Logger) {
	g := graceful.New(
		r,
		graceful.WithAddr(
			fmt.Sprintf("%s:%d", v.GetString(config.KeyHost), v.GetInt(config.KeyPort)),
		),
		graceful.WithReadTimeout(v.GetDuration(config.KeyReadTimeout)),
		graceful.WithWriteTimeout(v.GetDuration(config.KeyWriteTimeout)),
		graceful.WithLogger(l),
		graceful.WithCleanup(func() {
			log.Info("gracefulRunServer: test cleanup...")
			time.Sleep(5 * time.Second)
		}),
	)

	g.Run()
	// Wait for send event to Sentry when server start failed
	time.Sleep(3 * time.Second)
}
