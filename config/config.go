package config

import (
	"fmt"
	"os"

	log "github.com/litsea/log-slog"
	"github.com/spf13/viper"

	"github.com/litsea/gin-example/version"
)

func setDefault(v *viper.Viper) {
	v.SetDefault(KeyHost, "0.0.0.0")
	v.SetDefault(KeyPort, 8080)
	v.SetDefault(KeyCORSAllowOrigins, []string{"*"})
}

func ReadConfig(v *viper.Viper, cfgFile, configPath string) error {
	if cfgFile != "" {
		v.SetConfigFile(cfgFile)
	} else {
		v.SetConfigName("app")
		v.AddConfigPath(configPath)
	}

	err := v.ReadInConfig()
	if err != nil {
		return fmt.Errorf("config.ReadConfig: %w", err)
	}

	setDefault(v)

	return nil
}

func InitLogger(v *viper.Viper) {
	logCfg := v.Sub(KeyLog)
	logCfg.Set("rev", version.GitRev)

	err := log.Set(logCfg)
	if err != nil {
		fmt.Println("failed to init logger: ", err)
		os.Exit(1)
	}

	log.Info("logger initialized")
}
