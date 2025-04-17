package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/litsea/log-slog"
	vp "github.com/litsea/viper-aws"
	"github.com/litsea/viper-aws/secrets"
	"github.com/spf13/viper"
)

var config *Config

type Config struct {
	v *viper.Viper
	p *secrets.Provider

	MWCorsFn gin.HandlerFunc
}

func Init(cfgFile, cfgType string) {
	var (
		err error
		p   *secrets.Provider
		cfg *vp.Config
	)

	l := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
	v := viper.NewWithOptions(viper.WithLogger(l))
	v.SetEnvPrefix("GIN_EXP")
	v.AutomaticEnv()
	sid := os.Getenv("CONFIG_AWS_SECRET_ID")

	if sid == "" {
		cfg, err = vp.NewFile(v,
			vp.WithLogger(l),
			vp.WithType(cfgType),
			vp.WithFile(cfgFile),
			vp.WithOnFileChange(onFileChange(v)),
			vp.WithSetDefaultFunc(setDefault),
		)
	} else {
		cfg, err = vp.NewSecrets(v, sid,
			[]vp.Option{
				vp.WithLogger(l),
				vp.WithType(cfgType),
				vp.WithSetDefaultFunc(setDefault),
			},
			[]secrets.Option{
				secrets.WithLogger(l),
				secrets.WithOnChangeFunc(onSecretsChange(v, cfgType)),
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

	// Update logger
	*l = *log.GetSlog()
	// Update CORS middleware config
	config.setMiddlewareCorsFn()
}

func Get() *Config {
	return config
}

func V() *viper.Viper {
	return config.v
}
