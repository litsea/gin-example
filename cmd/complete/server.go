package complete

import (
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/litsea/log-slog"
	"github.com/spf13/viper"

	"github.com/litsea/gin-example/config"
)

func newServer(v *viper.Viper) error {
	r := gin.New()

	addMiddleware(r, v, log.Get())
	newRouter(r, v)

	addr := fmt.Sprintf("%s:%d", v.GetString(config.KeyHost), v.GetInt(config.KeyPort))
	if err := r.Run(addr); err != nil {
		return fmt.Errorf("failed to start server %s: %w", addr, err)
	}

	return nil
}
