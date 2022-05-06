package main

import (
	"github.com/spf13/cobra"
	"gitlab.com/RehakFrantisek/rehak_clc/ctcgrpc/cmd/client"
	"gitlab.com/RehakFrantisek/rehak_clc/ctcgrpc/cmd/server"
	"gitlab.com/RehakFrantisek/rehak_clc/ctcgrpc/pkg/util"
)

func main() {
	cmd := &cobra.Command{
		Use: "ctcgrpc",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	cmd.AddCommand(server.Cmd(), client.Cmd())

	util.ExitOnError(cmd.Execute())
}
