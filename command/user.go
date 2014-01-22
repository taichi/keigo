package command

import (
	"github.com/spf13/cobra"
	. "github.com/taichi/keigo/core"
)

func addUserCommands(parent *cobra.Command) {
	parent.AddCommand(newPasswdCmd())
	parent.AddCommand(newCKIDCmd())
	parent.AddCommand(newLGPWCmd())
	parent.AddCommand(newPWSTCmd())
}

func newPasswdCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "passwd",
		Short: "Enable Authentication",
		Long: `Enable Authentication
Example:
    keigo passwd newpassword
`,
		Run: newRunFn(func(cmd *cobra.Command, args []string) {
			if 0 < len(args) {
				execute(&CKID{On}, &PWST{On}, &LGPW{args[0]})
				cmd.Println("Authentication Enabled And set a new password.")
			} else {
				cmd.Println("new password required")
			}
		}),
	}
}

func newCKIDCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "CKID",
		Short: "Enable/Disable Authentication Or Read Current State",
		Long: `Enable/Disable Authentication Or Read Current State
Example:
    keigo CKID
	keigo CKID Enable`,
		Run: newRunFn(func(cmd *cobra.Command, args []string) {
			var results string
			if 0 < len(args) {
				if st := toState(args[0]); st != nil {
					results = execute(
						&CKID{st},
						&CKID{})[1]
				} else {
					cmd.Printf("Invalid State value %s\n", args[0])
				}
			} else {
				results = execute(&CKID{})[0]
			}
			cmd.Println(results)
		}),
	}
}
func newLGPWCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "LGPW",
		Short: "Update authentication token",
		Long: `
Example:
    keigo LGPW isa
`,
		Run: newRunFn(func(cmd *cobra.Command, args []string) {
			if 0 < len(args) {
				results := execute(&LGPW{Password: args[0]})
				cmd.Println(results)
			} else {
				cmd.Println("new password required")
			}
		}),
	}
}
func newPWSTCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "PWST",
		Short: "Enable/Disable Password Authentication Or Read Current State",
		Long: `Enable/Disable Password Authentication Or Read Current State
Example:
    keigo PWST
	keigo PWST Enable`,
		Run: newRunFn(func(cmd *cobra.Command, args []string) {
			var results string
			if 0 < len(args) {
				if st := toState(args[0]); st != nil {
					results = execute(
						&PWST{st},
						&PWST{})[1]
				} else {
					cmd.Printf("Invalid State value %s\n", args[0])
				}
			} else {
				results = execute(&PWST{})[0]
			}
			cmd.Println(results)
		}),
	}
}
