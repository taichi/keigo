package command

import (
	"github.com/spf13/cobra"
	. "github.com/taichi/keigo/core"
	"github.com/taichi/keigo/log"
)

func addNetworkCommands(parent *cobra.Command) {
	parent.AddCommand(newCKIPCmd())
	parent.AddCommand(newCKSTCmd())
}

func newCKIPCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "CKIP",
		Short: "Enable/Disable Network Watching or read current state",
		Long: `E: Enable Network Watching
D: Disable Network Watching
X: Keep Current State

Example:
    keigo CKIP
    keigo CKIP EDEXXXXX`,
		Run: newRunFn(func(cmd *cobra.Command, args []string) {
			var results string
			if 0 < len(args) {
				r := execute(&CKIP{*toStates(20, args[0])}, &CKIP{})
				results = r[1]
			} else {
				r := execute(&CKIP{})
				results = r[0]
			}
			labels := map[rune]string{
				'D': "Disable",
				'E': "Enable [Input Off]",
				'F': "Enable [Input On]",
			}
			log.Debugf("RAW OUTPUT : %s", results)

			cmd.Println("Current Network Watching Status")
			// TODO support json output
			// TODO support colorized output
			for i, ch := range results {
				cmd.Println("-------------------")
				cmd.Printf("[%2d] : %s\n", i+1, labels[ch])
			}
		}),
	}
}

func newCKSTCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "CKST",
		Short: "read SNMP trap status",
		Long:  `read SNMP trap status`,
		Run: newRunFn(func(cmd *cobra.Command, args []string) {
			results := execute(CKST)[0]

			labels := map[rune]string{
				'D': "Disable",
				'E': "Enable",
			}
			log.Debugf("RAW OUTPUT : %s", results)
			cmd.Println("Current SNMP Watching Status")
			for i, ch := range results {
				cmd.Println("-------------------")
				cmd.Printf("[%2d] : %s\n", i+1, labels[ch])
			}
		}),
	}
}
