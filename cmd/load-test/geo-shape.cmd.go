package load_test

import (
	cmd_router "digikalajet/internal/cmd-router"
	"github.com/spf13/cobra"
)

var (
	GeoShapeV1Cmd = &cobra.Command{
		Use:   "geo-shape",
		Short: "serve controller service",
		Run:   geoShapeRun,
	}
)

func geoShapeRun(cmd *cobra.Command, args []string) {
	cmd_router.Route(args[0])
}
