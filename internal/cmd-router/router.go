package cmd_router

import (
	"digikalajet/internal/geo-shaps/api/cmd-router"
	cli_router "digikalajet/pkg/cli-router"
	"digikalajet/utils/elasticsearch"
)

func Route(args string) {
	elasticsearch.E()
	geoShapeController := cmd_router.NewGeoShapeController()
	scenarioGeoShapeController := cmd_router.NewScenarioGeoShapeController()

	cli_router.AddRoute("make-index", geoShapeController.MakeIndex)
	cli_router.AddRoute("make-fake-polygons", geoShapeController.MakeFakePolygons)
	cli_router.AddRoute("insert-fake-polygon-to-elastic", geoShapeController.InsertFakePolygonsToElastic)
	cli_router.AddRoute("load-test", geoShapeController.LoadTest)
	cli_router.AddRoute("scenario-load-test", scenarioGeoShapeController.ScenarioLoadTest)

	cli_router.Run(args)
}
