package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/litsea/gin-example/cmd/complete"
	"github.com/litsea/gin-example/config"
)

var (
	cfgFile string
	cfgType string
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
	rootCmd.PersistentFlags().StringVar(&cfgFile,
		"conf-file", "./app.yaml", "config file (default is ./app.yaml)")
	rootCmd.PersistentFlags().StringVar(&cfgType,
		"conf-type", "yaml", "config type (default is yaml)")

	cobra.OnInitialize(func() {
		config.Init(cfgFile, cfgType)
	})

	rootCmd.AddCommand(complete.New())
}
