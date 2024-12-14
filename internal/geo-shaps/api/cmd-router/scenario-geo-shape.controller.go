package cmd_router

import (
	"context"
	"digikalajet/internal/geo-shaps/core/usecase"
)

type ScenarioGeoShapeController struct {
	scenarioGeoShapeUsecase *usecase.ScenarioGeoShapeUsecase
}

func NewScenarioGeoShapeController() *ScenarioGeoShapeController {
	return &ScenarioGeoShapeController{
		scenarioGeoShapeUsecase: usecase.NewScenarioGeoShapeUsecase(),
	}
}

func (c ScenarioGeoShapeController) ScenarioLoadTest(ctx context.Context) {
	c.scenarioGeoShapeUsecase.ScenarioLoadTest(ctx)

}
