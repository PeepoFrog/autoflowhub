package cli

import (
	"github.com/spf13/cobra"
)

const shortDescription = "Autoflowhub, infra automation tool"
const logngDescription = "Autoflowhub is the toool which helps dev to automate, maintain and test things"
const usage = "autoflowhub"

func NewCLI(cmds []*cobra.Command) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   usage,
		Short: shortDescription,
		Long:  logngDescription,
	}
	for _, cmd := range cmds {
		rootCmd.AddCommand(cmd)
	}
	return rootCmd
}
