package command

import (
	"github.com/spf13/cobra"
	. "github.com/taichi/keigo/core"
	"github.com/taichi/keigo/log"
)

func addDirectInputCommnads(parent *cobra.Command) {
	parent.AddCommand(newCKDICmd())
	parent.AddCommand(newROPSCmd())
}

func newCKDICmd() *cobra.Command {
	return &cobra.Command{
		Use:   "CKDI",
		Short: "Enable/Disable direct input and read current state",
		Long: `Enable/Disable direct input and read current state
E: Enable Direct Input
D: Disable Direct Input
X: Keep Current State

Example:
    keigo CKDI
    keigo CKDI EDEX`,
		Run: newRunFn(func(cmd *cobra.Command, args []string) {
			var results string
			if 0 < len(args) {
				r := execute(&CKDI{*toStates(4, args[0])}, &CKDI{})
				results = r[1]
			} else {
				r := execute(&CKDI{})
				results = r[0]
			}
			labels := map[rune]string{
				'D': "Disable",
				'E': "Enable [Input Off]",
				'F': "Enable [Input On]",
			}
			log.Debugf("RAW OUTPUT : %s", results)

			cmd.Println("Current Direct Input Status")
			// TODO support json output
			// TODO support colorized output
			for i, ch := range results {
				cmd.Println("-------------------")
				cmd.Printf("DI1[%d] : %s\n", i+1, labels[ch])
			}
		}),
	}
}

func newROPSCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "ROPS",
		Short: "read direct input state",
		Long:  `read direct input state`,
		Run: newRunFn(func(cmd *cobra.Command, args []string) {
			results := execute(ROPS)
			cmd.Println("Current Direct Input Status")
			labels := map[rune]string{
				'0': "Break",
				'1': "Make",
			}
			log.Debugf("RAW OUTPUT : %s", results[0])
			for i, ch := range results[0] {
				cmd.Println("-------------------")
				cmd.Printf("INPUT[%d] : %s\n", i+1, labels[ch])
			}
		}),
	}
}
