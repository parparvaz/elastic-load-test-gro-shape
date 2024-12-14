package fake_polygon

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

type FakePolygon struct {
	Polygons [][][]float64
}

const (
	minLat = -90.0
	maxLat = 90.0
	minLng = -90.0
	maxLng = 90.0
)

func Make(count int) {
	rand.NewSource(time.Now().UnixNano())
	log.Println(count)

	radius := 4.5

	fakePolygons := generateFakePolygons(count, radius)

	file, err := os.Create("fake_polygons.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer func() {
		_ = file.Close()
	}()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(fakePolygons)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Printf("%d fake polygons with random centers within Tehran have been generated and saved to fake_polygons.json. \n", count)
}

type Polygon struct {
	Coordinates [][]float64 `json:"coordinates"`
}

func randomLocationInRange(minLat, maxLat, minLon, maxLon float64) []float64 {
	lat := rand.Float64()*(maxLat-minLat) + minLat
	lon := rand.Float64()*(maxLon-minLon) + minLon
	return []float64{lon, lat}
}

func randomPointInCircle(centerLat, centerLon, radius float64) []float64 {
	distance := rand.Float64() * radius
	angle := rand.Float64() * 2 * math.Pi

	deltaLat := distance * math.Cos(angle)
	deltaLon := distance * math.Sin(angle)

	newLat := centerLat + (deltaLat / 111.32) // هر درجه عرض جغرافیایی تقریبا 111.32 کیلومتر
	newLon := centerLon + (deltaLon / (111.32 * math.Cos(centerLat*math.Pi/180)))

	return []float64{newLon, newLat}
}

func generateRandomPolygon(centerLat, centerLon, radius float64) Polygon {
	numPoints := rand.Intn(3) + 3 // حداقل 3 نقطه در هر پلیگان
	var coordinates [][]float64
	for i := 0; i < numPoints; i++ {
		point := randomPointInCircle(centerLat, centerLon, radius)
		coordinates = append(coordinates, point)
	}
	coordinates = append(coordinates, coordinates[0])
	return Polygon{Coordinates: coordinates}
}

func generateFakePolygons(num int, radius float64) []Polygon {
	var polygons []Polygon
	usedCenters := make(map[string]bool)

	for len(polygons) < num {
		center := randomLocationInRange(minLat, maxLat, minLng, maxLng)
		centerLat := center[1]
		centerLon := center[0]

		key := fmt.Sprintf("%f,%f", centerLat, centerLon)
		if usedCenters[key] {
			continue
		}
		usedCenters[key] = true

		polygon := generateRandomPolygon(centerLat, centerLon, radius)
		polygons = append(polygons, polygon)
	}

	return polygons
}

func GetPolygons() ([]Polygon, error) {
	file, err := os.Open("fake_polygons.json")
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = file.Close()
	}()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var polygon []Polygon
	if err := json.Unmarshal(data, &polygon); err != nil {
		return nil, err
	}

	return polygon, nil
}

func randomFloatInRange(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomLatLngInTehran() (float64, float64) {
	lat := randomFloatInRange(minLat, maxLat)
	lng := randomFloatInRange(minLng, maxLng)
	return lat, lng
}

func GenerateRandomLatLng() []float64 {
	rand.NewSource(time.Now().UnixNano())
	lat, lng := randomLatLngInTehran()
	return []float64{lat, lng}

}
