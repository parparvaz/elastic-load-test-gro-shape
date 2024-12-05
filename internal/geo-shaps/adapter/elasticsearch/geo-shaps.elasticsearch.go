package elasticsearch

import (
	"bytes"
	"context"
	"digikalajet/internal/geo-shaps/core/domain"
	"digikalajet/utils/elasticsearch"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

type GeoShapeElasticSearch struct {
	conn *elasticsearch.ElasticSearch
}

func NewGeoShapeElasticSearch() *GeoShapeElasticSearch {
	return &GeoShapeElasticSearch{
		conn: elasticsearch.E(),
	}
}

const (
	GeoShapeV1IndexName    string = "geo-shaps-v1"
	GeoShapeV1IndexMapping string = `{
  "mappings": {
    "geo-shaps-v1-mappings" : {
		"properties": {
		  "location": {
			"type": "geo_shape"
		  },
		  "shop_id":{
			"type": "integer"
		  },
		  "polygon_id":{
			"type": "integer"
		  },
		  "radius_base":{
			"type": "boolean"
		  }
		}
	  }
	}
}`
)

func (u *GeoShapeElasticSearch) MakeGeoShapeV1Index(ctx context.Context) {
	res, err := u.conn.Client.Indices.Create(
		GeoShapeV1IndexName,
		u.conn.Client.Indices.Create.WithBody(strings.NewReader(GeoShapeV1IndexMapping)),
		u.conn.Client.Indices.Create.WithContext(ctx),
	)
	if err != nil {
		log.Fatalf("Error creating the index: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error response: %s", res.String())
	} else {
		fmt.Printf("Index %s created successfully\n", GeoShapeV1IndexName)
	}
}

func (u *GeoShapeElasticSearch) FindByQuery(ctx context.Context, domain elasticsearch.Search) error {
	data, err := json.Marshal(domain)
	if err != nil {
		return err
	}

	_, err = u.conn.Client.Search(
		u.conn.Client.Search.WithIndex(GeoShapeV1IndexName),
		u.conn.Client.Search.WithBody(strings.NewReader(string(data))),
		u.conn.Client.Search.WithSize(1),
		u.conn.Client.Search.WithPretty(),
	)
	return err
}

func (u *GeoShapeElasticSearch) InsertGeoShape(ctx context.Context, docs []domain.GeoShapeV1Index) error {
	var bulkBody bytes.Buffer
	log.Println(len(docs))
	for _, doc := range docs {
		meta := fmt.Sprintf(`{ "index": { "_index": "%s", "_type": "geo-shaps-v1-mappings" } }`, GeoShapeV1IndexName)
		bulkBody.WriteString(meta + "\n")

		docJSON, err := json.Marshal(doc)
		if err != nil {
			log.Fatalf("Error marshaling document: %s", err)
		}
		bulkBody.Write(docJSON)
		bulkBody.WriteString("\n")
	}

	res, err := u.conn.Client.Bulk(
		bytes.NewReader(bulkBody.Bytes()),
		u.conn.Client.Bulk.WithContext(ctx),
	)
	if err != nil {
		log.Fatalf("Error performing bulk operation: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Bulk request error: %s", res.String())
	} else {
		fmt.Println("Bulk data inserted successfully!")
	}

	return nil
}

func (u *GeoShapeElasticSearch) CheckElasticsearchStatus(context.Context) error {
	res, err := u.conn.Client.Cluster.Health(
		u.conn.Client.Cluster.Health.WithContext(context.Background()),
		u.conn.Client.Cluster.Health.WithLevel("cluster"),
	)
	if err != nil {
		return fmt.Errorf("error checking cluster health: %s", err)
	}
	defer res.Body.Close()

	// نمایش نتیجه وضعیت کلستر
	if res.IsError() {
		return fmt.Errorf("error response from Elasticsearch: %s", res.String())
	}

	// نمایش وضعیت کلستر
	fmt.Println("Cluster health:")
	fmt.Println(res)
	return nil
}

func (u *GeoShapeElasticSearch) CheckElasticsearchIndices(context.Context) error {
	res, err := u.conn.Client.Cat.Indices(
		u.conn.Client.Cat.Indices.WithContext(context.Background()),
		u.conn.Client.Cat.Indices.WithFormat("json"),
	)
	if err != nil {
		return fmt.Errorf("error checking indices: %s", err)
	}
	defer res.Body.Close()

	// نمایش نتیجه وضعیت ایندکس‌ها
	if res.IsError() {
		return fmt.Errorf("error response from Elasticsearch: %s", res.String())
	}

	// نمایش وضعیت ایندکس‌ها
	fmt.Println("Indices status:")
	fmt.Println(res)
	return nil
}

func (u *GeoShapeElasticSearch) GetSystemResourceUsage(context.Context) (elasticsearch.ClusterNodeStats, error) {
	var stats elasticsearch.ClusterNodeStats

	res, err := u.conn.Client.Nodes.Stats(
		u.conn.Client.Nodes.Stats.WithContext(context.Background()),
		u.conn.Client.Nodes.Stats.WithMetric("os", "process", "jvm", "fs", "thread_pool", "transport"),
	)
	if err != nil {
		return elasticsearch.ClusterNodeStats{}, err
	}

	defer res.Body.Close()

	if res.IsError() {
		return elasticsearch.ClusterNodeStats{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return elasticsearch.ClusterNodeStats{}, err
	}

	if err = json.Unmarshal(body, &stats); err != nil {
		return elasticsearch.ClusterNodeStats{}, err
	}
	return stats, nil
}
