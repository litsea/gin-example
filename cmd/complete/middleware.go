package complete

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	api "github.com/litsea/gin-api"
	"github.com/litsea/gin-api/cors"
	log "github.com/litsea/gin-api/log"
	g18n "github.com/litsea/gin-i18n"
	"github.com/litsea/i18n"
	"github.com/spf13/viper"
	"golang.org/x/text/language"

	"github.com/litsea/gin-example/assets"
	"github.com/litsea/gin-example/config"
)

func addMiddleware(r *gin.Engine, v *viper.Viper, sl *slog.Logger) {
	// logger for gin-api and gin-i18n
	l := log.New(
		sl,
		log.WithRequestHeader(true),
		log.WithRequestBody(true),
		log.WithUserAgent(true),
		log.WithStackTrace(true),
	)

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
	)
}
