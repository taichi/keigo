package command

import (
	"github.com/spf13/cobra"
)

var version = &cobra.Command{
	Use:   "version",
	Short: "Print the Version No of Keigo",
	Long:  `All software has versions. This is Keigo's`,
	Run: newRunFn(func(cmd *cobra.Command, args []string) {
		cmd.Printf("Keigo     : %s\n", "0.0.1")
	}),
}
