package complete

import (
	"github.com/spf13/cobra"

	"github.com/litsea/gin-example/config"
)

var cmd *cobra.Command

func New() *cobra.Command {
	cmd = &cobra.Command{
		Use:   "complete",
		Short: "complete gin",
		Run: func(cmd *cobra.Command, args []string) {
			newServer(config.V())
		},
	}

	return cmd
}
