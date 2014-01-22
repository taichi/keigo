package command

import (
	"github.com/spf13/cobra"
	"github.com/taichi/keigo/core"
	"github.com/taichi/keigo/log"
	"github.com/taichi/keigo/util"
)

var configpath string
var verbose bool
var keigoCmd = &cobra.Command{Use: "keigo"}

func init() {
	keigoCmd.PersistentFlags().StringVarP(&configpath, "config", "c", "config.toml", "config file path")
	keigoCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "print debug informations")
}

func Execute() {
	addCommands()
	keigoCmd.Execute()
}

func addCommands() {
	keigoCmd.AddCommand(version)
	keigoCmd.AddCommand(info)
	addLampCommands(keigoCmd)
	addNetworkCommands(keigoCmd)
	addDirectInputCommnads(keigoCmd)
	addSoundCommands(keigoCmd)
	addUserCommands(keigoCmd)
}

type runFn func(cmd *cobra.Command, args []string)

func newRunFn(f runFn) runFn {
	return func(cmd *cobra.Command, args []string) {
		if verbose {
			log.VerboseLog()
		}
		f(cmd, args)
	}
}

func execute(cmds ...core.Command) []string {
	config, conferr := core.LoadConfig(configpath)
	util.MaybeFault(conferr)
	session, conerr := core.Connect(config)
	util.MaybeFault(conerr)
	defer session.Close()
	results := make([]string, len(cmds))
	var err error
	for i, c := range cmds {
		results[i], err = session.Execute(c)
		log.Debugf("< %s %s", c, results[i])
		util.MaybeFault(err)
	}
	return results
}
