package command

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	. "github.com/taichi/keigo/core"
	"github.com/taichi/keigo/log"
	"github.com/taichi/keigo/util"
	"strings"
)

var wait, time int

type desc struct {
	Use, Label string
	Index      int
}

func addLampCommands(parent *cobra.Command) {
	for _, v := range []desc{
		{"red", "red lamp", 0},
		{"yellow", "yellow lamp", 1},
		{"green", "green lamp", 2},
		{"buzzer1", "buzzer1", 3},
		{"buzzer2", "buzzer2", 4},
	} {
		parent.AddCommand(newLampCmd(v))
	}
	parent.AddCommand(newACOPCmd())
	parent.AddCommand(newALOFCmd("off"))
	parent.AddCommand(newALOFCmd("ALOF"))
}

var longTemplate = util.Must("Long", `turn on/off {{.Label}}.
Example:
    keigo {{.Use}} off
    keigo {{.Use}} on
    keigo {{.Use}} blink
    keigo {{.Use}} quickblink
    keigo {{.Use}} on -t 3 -w 2`)

func newLampCmd(d desc) *cobra.Command {
	newCmd := &cobra.Command{
		Use:   d.Use,
		Short: fmt.Sprintf("turn on/off %s", d.Label),
		Long:  longTemplate.Do(d),
		Run: newRunFn(func(cmd *cobra.Command, args []string) {
			var state = On

			if 0 < len(args) {
				states := map[string]State{
					"off":        Off,
					"on":         On,
					"blink":      Blink,
					"quickblink": QuickBlink,
					"0":          Off,
					"1":          On,
					"2":          Blink,
					"3":          QuickBlink,
				}
				state = states[strings.ToLower(args[0])]
			}

			if state != nil {
				st := *NewStates(8)
				st[d.Index] = state
				order := &ACOP{
					LampOrBuzzer,
					Seconds(wait),
					Seconds(time),
					st,
				}
				results := execute(order, &ACOP{})
				printResults(cmd, &results)
			} else {
				cmd.Printf("lamp status supports off/on/blink/quickblink")
			}
		}),
	}
	addTimeFlags(newCmd)

	return newCmd
}

func addTimeFlags(cmd *cobra.Command) {
	cmd.Flags().IntVarP(&wait, "wait", "w", 0, "waiting for specific seconds.")
	cmd.Flags().IntVarP(&time, "time", "t", 0, "turn lamp off after specific seconds.")
}

func printResults(cmd *cobra.Command, results *[]string) {
	// TODO support json output
	// TODO support colorized output
	currentState := (*results)[1]
	log.Debugf("RAW OUTPUT : %s", currentState)
	cmd.Println("Current Lamp Status")
	statusLabels := map[uint8]string{
		'0': "off",
		'1': "on",
		'2': "blink",
		'3': "quickblink",
	}

	for i, l := range []string{"RED", "YELLOW", "GREEN", "BUZZER1", "BUZZER2"} {
		cmd.Println("-------------------")
		cmd.Printf("%-10s : %6s\n", l, statusLabels[currentState[i]])
	}
}

func newACOPCmd() *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "ACOP",
		Short: "execute ACOP command",
		Long: `execute low level ACOP command
Example:
    keigo ACOP x12
    keigo ACOP –u 1 x12`,
	}

	addTimeFlags(newCmd)
	u := newCmd.Flags().IntP("unit", "u", 0, "output to specific unit")
	newCmd.Run = newRunFn(func(cmd *cobra.Command, args []string) {
		unit := *u
		var buffer bytes.Buffer
		fmt.Fprintf(&buffer, "ACOP ")
		if 0 < unit {
			fmt.Fprintf(&buffer, "-u %d ", unit)
		}

		if 0 < len(args) {
			fmt.Fprintf(&buffer, "%s ", args[0])
		}

		if 0 < wait {
			fmt.Fprintf(&buffer, "-w %d ", wait)
		}
		if 0 < time {
			fmt.Fprintf(&buffer, "-t %d", time)
		}

		results := execute(&buffer, &ACOP{})
		printResults(cmd, &results)
	})

	return newCmd
}

var alofTemplate = util.Must("alof",
	`turn off all lamp and buzzer.
Example:
    keigo {{.Use}}`)

func newALOFCmd(use string) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: "turn off all lamp and buzzer",
		Long:  alofTemplate.Do(map[string]string{"Use": use}),
		Run: newRunFn(func(cmd *cobra.Command, args []string) {
			results := execute(ALOF, &ACOP{})
			printResults(cmd, &results)
		}),
	}
}

// TODO RLY1 - RLY8 support
