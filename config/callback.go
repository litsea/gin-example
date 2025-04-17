package config

import (
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/fsnotify/fsnotify"
	"github.com/litsea/gin-api/cors"
	log "github.com/litsea/log-slog"
	"github.com/spf13/viper"
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

func onFileChange(_ *viper.Viper) func(evt fsnotify.Event) {
	return func(evt fsnotify.Event) {
		log.Warn("config.onFileChange", "name", evt.Name, "op", evt.Op.String())
		config.setMiddlewareCorsFn()

		// Remove comment for testing
		// panic("config.onFileChange: panic test")
	}
}

func onSecretsChange(_ *viper.Viper, cfgType string) func(out *secretsmanager.GetSecretValueOutput) {
	return func(out *secretsmanager.GetSecretValueOutput) {
		// Check syntax
		vv := viper.New()
		vv.SetConfigType(cfgType)
		err := vv.ReadConfig(strings.NewReader(*out.SecretString))
		if err != nil {
			log.Error("config.onSecretsChange: invalid value", "version", *out.VersionId,
				"createdDate", out.CreatedDate, "err", err)
			return
		}

		log.Warn("config.onSecretsChange: updated", "version", *out.VersionId,
			"createdDate", out.CreatedDate)

		config.setMiddlewareCorsFn()

		// Remove comment for testing
		// panic("config.onSecretsChange: panic test")
	}
}

func QuitWatch() {
	if config != nil && config.p != nil {
		config.p.QuitWatch()
	}
}

func (c *Config) setMiddlewareCorsFn() {
	c.MWCorsFn = cors.New(cors.WithAllowOrigin(c.v.GetStringSlice(KeyCORSAllowOrigins)))
}
