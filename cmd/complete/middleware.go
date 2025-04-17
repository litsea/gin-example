package complete

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	api "github.com/litsea/gin-api"
	"github.com/litsea/gin-api/errcode"
	"github.com/litsea/gin-api/log"
	"github.com/litsea/gin-api/ratelimit"
	g18n "github.com/litsea/gin-i18n"
	"github.com/litsea/i18n"
	"github.com/spf13/viper"
	timeout "github.com/vearne/gin-timeout"
	"golang.org/x/text/language"

	"github.com/litsea/gin-example/assets"
	"github.com/litsea/gin-example/config"
)

var IpLimiter = ratelimit.NewLimiter(10, time.Minute)

func addMiddleware(r *gin.Engine, v *viper.Viper, l log.Logger) {
	// i18n
	gi := g18n.New(
		g18n.WithOptions(
			i18n.WithLanguages(language.English, language.Chinese),
			i18n.WithLoaders(
				i18n.EmbedLoader(api.Localize, "./localize/"),
				i18n.EmbedLoader(assets.Localize, "./localize/"),
			),
		),
		g18n.WithLogger(l),
	)

	mws := []gin.HandlerFunc{
		log.Middleware(l),
		api.Recovery(api.HandleRecovery()),
		gi.Localize(),
	}

	if config.Get().MWCorsFn != nil {
		mws = append(mws, func(c *gin.Context) {
			config.Get().MWCorsFn(c)
		})
	}

	mws = append(mws,
		// Note: the timeout middleware will cause the panic stack loss
		timeout.Timeout(
			// Request timeout cannot exceed server exceed timeout `config.KeyWriteTimeout`
			timeout.WithTimeout(v.GetDuration(config.KeyRequestTimeout)),
			timeout.WithContentType("application/json; charset=utf-8"),
			timeout.WithErrorHttpCode(http.StatusServiceUnavailable),
			// TODO: no translation yet
			timeout.WithDefaultMsg(errcode.ErrServiceTimeout),
			timeout.WithGinCtxCallBack(func(c *gin.Context) {
				l.ErrorRequest(c, "middleware: request timeout", map[string]any{
					"err": errcode.ErrServiceTimeout,
				})
			}),
		),
	)
	r.Use(mws...)
}
