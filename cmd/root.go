package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/litsea/gin-example/cmd/complete"
	"github.com/litsea/gin-example/config"
)

var (
	cfgFile string
	v       *viper.Viper
)

var ErrInvalidCommand = errors.New("invalid command")

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "example app",
	Long:  "example app",
	RunE: func(*cobra.Command, []string) error {
		return ErrInvalidCommand
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile,
		"config", "", "config file (default is ./app.yaml)")

	v = viper.GetViper()
	rootCmd.AddCommand(complete.New(v))
}

func initConfig() {
	if err := config.ReadConfig(v, cfgFile, "./"); err == nil {
		fmt.Printf("using config file: %s\n", viper.ConfigFileUsed())
		v.WatchConfig()
		config.InitLogger(v)
	} else {
		fmt.Println(err)
		os.Exit(1)
	}
}
