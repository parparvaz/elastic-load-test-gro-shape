package ports

import (
	"context"
	"digikalajet/internal/geo-shaps/core/domain"
	"digikalajet/utils/elasticsearch"
)

type GeoShapeAdapterInterface interface {
	MakeGeoShapeV1Index(ctx context.Context)
	FindByQuery(ctx context.Context, domain elasticsearch.Search) (interface{}, error)
	InsertGeoShape(context.Context, []domain.GeoShapeV1Index) error
	CheckElasticsearchStatus(context.Context) error
	CheckElasticsearchIndices(context.Context) error
	GetSystemResourceUsage(context.Context) (elasticsearch.ClusterNodeStats, error)
}
