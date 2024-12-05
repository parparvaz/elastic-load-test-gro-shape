package cmd

import (
	load_test "digikalajet/cmd/load-test"
	"digikalajet/utils/provider"
	"github.com/spf13/cobra"
)

func Execute() {
	err := cmd.Execute()
	if err != nil {
		return
	}
}

var cmd = cobra.Command{
	Use:  "micro",
	Long: "A service that will validate restful transactions and send them to stripe.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	provider.Provider()
	cmd.AddCommand(load_test.GeoShapeV1Cmd)
}
