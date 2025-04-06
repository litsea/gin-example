package complete

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "complete",
	Short: "complete gin",
	RunE: func(cmd *cobra.Command, args []string) error {
		return newServer()
	},
}
