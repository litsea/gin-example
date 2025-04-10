package complete

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmd *cobra.Command

func New(v *viper.Viper) *cobra.Command {
	cmd = &cobra.Command{
		Use:   "complete",
		Short: "complete gin",
		RunE: func(cmd *cobra.Command, args []string) error {
			return newServer(v)
		},
	}

	return cmd
}
