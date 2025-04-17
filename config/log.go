package config

import (
	"fmt"

	log "github.com/litsea/log-slog"
	"github.com/spf13/viper"

	"github.com/litsea/gin-example/version"
)

func InitLogger(v *viper.Viper) error {
	logCfg := v.Sub(KeyLog)
	if logCfg == nil {
		return fmt.Errorf("config.InitLogger: empty logger config")
	}

	err := log.Set(logCfg, log.WithVersion(version.Version), log.WithGitRev(version.GitRev))
	if err != nil {
		return fmt.Errorf("config.InitLogger: failed to init logger: %w", err)
	}

	log.Info("config.InitLogger: logger initialized")
	return nil
}
