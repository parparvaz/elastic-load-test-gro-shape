package elasticsearch

type Search struct {
	Query SearchQuery `json:"query"`
}

type SearchQuery struct {
	GeoShape SearchQueryGeoShape `json:"geo_shape"`
}

type SearchQueryGeoShape struct {
	Location SearchQueryGeoShapeLocation `json:"location"`
}

type SearchQueryGeoShapeLocation struct {
	Shape    SearchQueryGeoShapeLocationShape `json:"shape"`
	Relation string                           `json:"relation"`
}

type SearchQueryGeoShapeLocationShape struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}
