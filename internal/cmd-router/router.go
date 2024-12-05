package cmd_router

import (
	"digikalajet/internal/geo-shaps/api/cmd-router"
	cli_router "digikalajet/pkg/cli-router"
)

func Route(args string) {
	geoShapeController := cmd_router.NewGeoShapeController()

	cli_router.AddRoute("make-index", geoShapeController.MakeIndex)
	cli_router.AddRoute("make-fake-polygons", geoShapeController.MakeFakePolygons)
	cli_router.AddRoute("insert-fake-polygon-to-elastic", geoShapeController.InsertFakePolygonsToElastic)
	cli_router.AddRoute("load-test", geoShapeController.LoadTest)

	cli_router.Run(args)
}
