package config

import (
	"fmt"
	"os"
	"time"

	log "github.com/litsea/log-slog"
	"github.com/spf13/viper"

	"github.com/litsea/gin-example/version"
)

func setDefault(v *viper.Viper) {
	v.SetDefault(KeyHost, "0.0.0.0")
	v.SetDefault(KeyPort, 8080)
	v.SetDefault(KeyReadTimeout, 15*time.Second)
	v.SetDefault(KeyWriteTimeout, 15*time.Second)
	v.SetDefault(KeyRequestTimeout, 15*time.Second)
	v.SetDefault(KeyMaxShutdownDuration, 30*time.Second)
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
	if logCfg == nil {
		fmt.Println("config.InitLogger: empty logger config")
		return
	}

	logCfg.Set("rev", version.GitRev)

	err := log.Set(logCfg)
	if err != nil {
		fmt.Println("config.InitLogger: failed to init logger: ", err)
		os.Exit(1)
	}

	log.Info("config.InitLogger: logger initialized")
}
