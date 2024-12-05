package cmd_router

import (
	"context"
	"digikalajet/internal/geo-shaps/core/domain"
	"digikalajet/internal/geo-shaps/core/usecase"
)

type GeoShapeController struct {
	geoShapeUsecase *usecase.GeoShapeUsecase
}

func NewGeoShapeController() *GeoShapeController {
	return &GeoShapeController{
		geoShapeUsecase: usecase.NewGeoShapeUsecase(),
	}
}

func (c *GeoShapeController) MakeIndex(ctx context.Context) {
	c.geoShapeUsecase.MakeIndex(ctx)

}

func (c *GeoShapeController) LoadTest(ctx context.Context) {
	c.geoShapeUsecase.LoadTest(ctx, domain.LoadTest{
		Counter: 10000,
	})
}

func (c *GeoShapeController) MakeFakePolygons(ctx context.Context) {
	c.geoShapeUsecase.MakeFakePolygons(ctx, domain.FakePolygon{
		Counter: 10000,
	})
}

func (c *GeoShapeController) InsertFakePolygonsToElastic(ctx context.Context) {
	c.geoShapeUsecase.InsertFakePolygonsToElastic(ctx)
}
