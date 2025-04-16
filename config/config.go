package config

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/fsnotify/fsnotify"
	log "github.com/litsea/log-slog"
	vp "github.com/litsea/viper-aws"
	"github.com/litsea/viper-aws/secrets"
	"github.com/spf13/viper"

	"github.com/litsea/gin-example/version"
)

var config *Config

type Config struct {
	v *viper.Viper
	p *secrets.Provider
}

func Init(cfgFile, cfgType string) {
	var (
		err error
		p   *secrets.Provider
		cfg *vp.Config
	)

	l := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
	v := viper.NewWithOptions(viper.WithLogger(l))
	sid := os.Getenv("CONFIG_AWS_SECRET_ID")

	if sid == "" {
		cfg, err = vp.NewFile(v,
			vp.WithType(cfgType),
			vp.WithFile(cfgFile),
			vp.WithOnFileChange(onFileChange),
			vp.WithSetDefaultFunc(setDefault),
		)
	} else {
		cfg, err = vp.NewSecrets(v, sid,
			[]vp.Option{vp.WithType(cfgType)},
			[]secrets.Option{
				secrets.WithLogger(l),
				secrets.WithKeepStages(10),
				secrets.WithWatchInterval(time.Second * 5),
				secrets.WithOnChangeFunc(onSecretsChange(cfgType)),
			},
		)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	config = &Config{v: cfg.V(), p: p}

	err = InitLogger(v)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	*l = *log.GetSlog()
}

func V() *viper.Viper {
	return config.v
}

func setDefault(v *viper.Viper) {
	v.SetDefault(KeyHost, "0.0.0.0")
	v.SetDefault(KeyPort, 8080)
	v.SetDefault(KeyReadTimeout, 15*time.Second)
	v.SetDefault(KeyWriteTimeout, 15*time.Second)
	v.SetDefault(KeyRequestTimeout, 15*time.Second)
	v.SetDefault(KeyMaxShutdownDuration, 30*time.Second)
	v.SetDefault(KeyCORSAllowOrigins, []string{"*"})
}

func onFileChange(evt fsnotify.Event) {
	log.Info("config.onFileChange", "evt", evt)
}

func onSecretsChange(cfgType string) func(out *secretsmanager.GetSecretValueOutput) {
	return func(out *secretsmanager.GetSecretValueOutput) {
		vv := viper.New()
		vv.SetConfigType(cfgType)
		err := vv.ReadConfig(strings.NewReader(*out.SecretString))
		if err != nil {
			log.Error("config.onChange: invalid value", "version", *out.VersionId,
				"createdDate", out.CreatedDate, "err", err)
			return
		}

		log.Info("config.onSecretsChange: updated", "version", *out.VersionId,
			"createdDate", out.CreatedDate)
	}
}

func QuitWatch() {
	if config != nil {
		if config.p != nil {
			config.p.QuitWatch()
		}
	}
}

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
