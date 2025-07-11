package complete

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/litsea/gin-api/graceful"
	apilog "github.com/litsea/gin-api/log"
	"github.com/litsea/kit/profiler"
	log "github.com/litsea/log-slog"
	"github.com/spf13/viper"

	"github.com/litsea/gin-example/config"
	"github.com/litsea/gin-example/util"
)

func newServer(v *viper.Viper) {
	// Enable the profiler only when the server address provided
	if v.GetString(config.KeyProfilerServerAddress) != "" {
		_, err := profiler.Start(
			// appName.serviceName.env
			"gin-example.complete."+v.GetString(config.KeyEnv),
			v.GetString(config.KeyProfilerServerAddress),
			profiler.WithAuth(
				v.GetString(config.KeyProfilerAuthUsername),
				v.GetString(config.KeyProfilerAuthPassword),
			),
			profiler.WithDebug(v.GetBool(config.KeyProfilerDebug)),
		)
		if err != nil {
			log.Error("complete.NewServer: failed to start profiler", "err", err)
		}
	}

	r := gin.New()

	// logger for gin-api, gin-i18n and other components/middlewares.
	l := apilog.New(
		log.GetSlog(),
		apilog.WithRequestHeader(true),
		apilog.WithRequestBody(true),
		apilog.WithUserAgent(true),
		apilog.WithStackTrace(true),
	)

	// For test only
	go func() {
		defer util.RecoverFn("complete.newServer")
		// Remove comment for testing
		// panic("new goroutine panic not in gin")
	}()

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
		graceful.WithStopTimeout(v.GetDuration(config.KeyStopTimeout)),
		graceful.WithLogger(l),
		graceful.WithCleanup(func() {
			log.Info("complete.gracefulRunServer: test cleanup...")
			config.QuitWatch()
			time.Sleep(2 * time.Second)
		}),
	)

	g.Run()
	// Wait for send event to Sentry when server start failed
	time.Sleep(3 * time.Second)
}
