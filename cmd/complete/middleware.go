package complete

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	api "github.com/litsea/gin-api"
	"github.com/litsea/gin-api/cors"
	"github.com/litsea/gin-api/errcode"
	log "github.com/litsea/gin-api/log"
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

	r.Use(
		log.Middleware(l),
		api.Recovery(api.HandleRecovery()),
		gi.Localize(),
		cors.New(cors.WithAllowOrigin(v.GetStringSlice(config.KeyCORSAllowOrigins))),
		timeout.Timeout(
			// Request timeout cannot exceed server exceed timeout `config.KeyWriteTimeout`
			timeout.WithTimeout(v.GetDuration(config.KeyRequestTimeout)),
			timeout.WithErrorHttpCode(http.StatusServiceUnavailable),
			// TODO: no translation and content-type
			timeout.WithDefaultMsg(errcode.ErrServiceUnavailable),
			timeout.WithGinCtxCallBack(func(c *gin.Context) {
				l.ErrorRequest(c, "middleware: request timeout2 ", map[string]any{
					"err": errcode.ErrServiceUnavailable,
				})
			}),
		),
	)
}
