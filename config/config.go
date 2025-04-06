package config

import (
	"fmt"
	"net/http"
	"os"
	"time"

	log "github.com/litsea/log-slog"
	"github.com/spf13/viper"

	"github.com/litsea/gin-example/version"
)

func setDefault() {
}

func ReadConfig(cfgFile, configPath string) error {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("app")
		viper.AddConfigPath(configPath)
	}

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("config.ReadConfig: %w", err)
	}

	setDefault()

	return nil
}

func InitLogger() {
	logCfg := viper.Sub("log")
	logCfg.Set("rev", version.GitRev)

	err := log.Set(logCfg)
	if err != nil {
		fmt.Println("failed to init logger: ", err)
		os.Exit(1)
	}

	log.Info("logger initialized")
}

func InitProfiler(host string, port int) {
	if port <= 0 {
		return
	}

	go func() {
		log.Info("pprof listening...", "host", host, "port", port)
		srv := &http.Server{
			Addr:              fmt.Sprintf("%s:%d", host, port),
			ReadHeaderTimeout: 30 * time.Second,
		}

		err := srv.ListenAndServe()
		if err != nil {
			log.Error("pprof listen failed", "err", err)

			return
		}
	}()
}
