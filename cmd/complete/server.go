package complete

import (
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/litsea/log-slog"
)

func newServer() error {
	r := gin.New()

	addMiddleware(r, log.Get())
	newRouter(r)

	if err := r.Run(":8080"); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
