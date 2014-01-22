package command

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	. "github.com/taichi/keigo/core"
	"regexp"
	"strconv"
)

func addSoundCommands(parent *cobra.Command) {
	parent.AddCommand(newPlayCmd())
	parent.AddCommand(newStopCmd())
	parent.AddCommand(newSPOPCmd())
}

var playTimes int
var isNumeric = regexp.MustCompile("^[0-9]{1,2}$")

func printSPOPResults(cmd *cobra.Command, results string) {
	state := func(s uint8) string {
		switch s {
		case '0':
			return "stopped"
		case '1':
			return "playing"
		}
		return ""
	}
	cmd.Printf("%-10s : %s\n", "State", state(results[0]))
	cmd.Printf("%-10s : %s\n", "SoundNo", results[1:3])
	cmd.Printf("%-10s : %t\n", "Repeat?", results[3] == '0')
	cmd.Printf("%-10s : %s\n", "PlayTimes", results[4:6])
}
func newPlayCmd() *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "play",
		Short: "Play Sound specific times",
		Long: `
Example:
    keigo play 1
	keigo play 10 -t 3
	keigo play 3 --times 7`,
		Run: newRunFn(func(cmd *cobra.Command, args []string) {
			if 0 < len(args) && isNumeric.MatchString(args[0]) {
				number, _ := strconv.Atoi(args[0])
				results := execute(
					&SPOP{On, SoundNo(number), Times(playTimes)},
					&SPOP{})[1]
				printSPOPResults(cmd, results)
			} else {
				cmd.Println("sound number is required.")
			}
		}),
	}
	newCmd.Flags().IntVarP(&playTimes, "times", "t", 0, "play sound specific times.")
	return newCmd
}

func newStopCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "stop",
		Short: "Stop sound",
		Long: `
Example:
    keigo stop`,
		Run: newRunFn(func(cmd *cobra.Command, args []string) {
			results := execute(&SPOP{SwitchTo: Off}, &SPOP{})[1]
			printSPOPResults(cmd, results)
		}),
	}
}

func newSPOPCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "SPOP",
		Short: "Enable/Disable Sound Control or read current state",
		Long: `
Example:
    keigo SPOP
    keigo SPOP 10100000`,
		Run: newRunFn(func(cmd *cobra.Command, args []string) {
			var results string
			if 0 < len(args) {
				var buff bytes.Buffer
				fmt.Fprintf(&buff, "SPOP %s", args[0])
				results = execute(&buff)[0]
			} else {
				results = execute(&SPOP{})[0]
			}
			cmd.Println(results)
		}),
	}
}
