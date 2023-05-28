package main

import (
	tool "github.com/mrlutik/autoflowhub/internal/cli"
	"github.com/mrlutik/autoflowhub/pkg/keygen/keygencmd"
	"github.com/mrlutik/autoflowhub/pkg/txsgen/txsgencmd"
	"github.com/spf13/cobra"
)

func main() {
	cmds := []*cobra.Command{txsgencmd.New(), keygencmd.New()}
	cli := tool.NewCLI(cmds)
	cobra.CheckErr(cli.Execute())
}
