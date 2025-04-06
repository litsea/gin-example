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
	cfgFile      string
	profilerHost string
	profilerPort int
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
	rootCmd.PersistentFlags().StringVar(&profilerHost,
		"profiler-host", "127.0.0.1", "profiler host")
	rootCmd.PersistentFlags().IntVar(&profilerPort,
		"profiler-port", 0, "profiler port")

	rootCmd.AddCommand(complete.Cmd)
}

func initConfig() {
	if err := config.ReadConfig(cfgFile, "./"); err == nil {
		fmt.Printf("using config file: %s\n", viper.ConfigFileUsed())
		viper.WatchConfig()
		config.InitLogger()
	} else {
		fmt.Println(err)
		os.Exit(1)
	}

	config.InitProfiler(profilerHost, profilerPort)
}
