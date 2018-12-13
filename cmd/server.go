package cmd

import (
	"log"

	"github.com/appno/stripe/schema"
	"github.com/appno/stripe/server"
	"github.com/spf13/cobra"
)

var serverExample = `
stripe server
stripe server 8082
`

// ServerCmd : Run Server
var ServerCmd = &cobra.Command{
	Use:     "server [PORT]",
	Short:   "Run server on port [PORT]",
	Args:    cobra.RangeArgs(0, 1),
	Example: serverExample,
	Run: func(cmd *cobra.Command, args []string) {
		port := schema.GetPort()
		if len(args) > 0 {
			port = args[0]
		}

		log.Fatal(server.Serve(port))
	},
}
