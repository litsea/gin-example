package complete

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	api "github.com/litsea/gin-api"
	"github.com/litsea/gin-api/cors"
	log "github.com/litsea/gin-api/log"
	g18n "github.com/litsea/gin-i18n"
	ginslog "github.com/litsea/gin-slog"
	"github.com/litsea/i18n"
	sloggin "github.com/samber/slog-gin"
	"golang.org/x/text/language"

	"github.com/litsea/gin-example/assets"
)

func addMiddleware(r *gin.Engine, l *slog.Logger) {
	// gin log for gin-api, gin-i18n
	gl := log.New(l, ginslog.AddCustomAttributes)

	// middleware for capture gin request/response
	gsl := ginslog.New(
		l,
		ginslog.WithFilters(
			sloggin.IgnorePath("/v1/health"),
			sloggin.IgnoreStatusLessThan(http.StatusInternalServerError),
		),
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
		g18n.WithLogger(gl),
	)

	r.Use(
		log.Middleware(gl),
		gsl,
		gi.Localize(),
		cors.New(cors.WithAllowOrigin([]string{"http://localhost:*"})),
	)
}
