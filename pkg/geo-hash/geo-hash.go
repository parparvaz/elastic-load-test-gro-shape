package geo_hash

const (
	geoHashChars     = "0123456789bcdefghjkmnpqrstuvwxyz"
	latitudeStep     = 0.0001 // گام عرض جغرافیایی
	longitudeStep    = 0.0001 // گام طول جغرافیایی
	geoHashPrecision = 8      // دقت GeoHash
)

type GeoHashGenerator struct {
	polygon   [][]float64
	geoHashes map[string]bool
}

// سازنده GeoHashGenerator
func NewGeoHashGenerator(polygon [][]float64) *GeoHashGenerator {
	return &GeoHashGenerator{
		polygon:   polygon,
		geoHashes: make(map[string]bool),
	}
}

// تولید GeoHashes برای پلی‌گون
func (g *GeoHashGenerator) GenerateGeoHashes() []string {
	var geoHashes []string
	// برای هر نقطه در grid، GeoHash تولید می‌شود و ذخیره می‌شود
	for _, point := range g.getGridPoints() {
		geoHash := g.encodeGeoHash(point[0], point[1], geoHashPrecision)
		if !g.geoHashes[geoHash] {
			g.geoHashes[geoHash] = true
			geoHashes = append(geoHashes, geoHash)
		}
	}
	return geoHashes
}

// دریافت نقاط grid داخل محدوده پلی‌گون
func (g *GeoHashGenerator) getGridPoints() [][]float64 {
	var minLat, maxLat, minLng, maxLng float64
	// تعیین min/max برای عرض و طول جغرافیایی پلی‌گون
	for i, point := range g.polygon {
		if i == 0 || point[0] < minLat {
			minLat = point[0]
		}
		if i == 0 || point[0] > maxLat {
			maxLat = point[0]
		}
		if i == 0 || point[1] < minLng {
			minLng = point[1]
		}
		if i == 0 || point[1] > maxLng {
			maxLng = point[1]
		}
	}

	// بررسی هر نقطه داخل محدوده و در صورت موجود بودن، بازگشت
	var points [][]float64
	for lat := minLat; lat <= maxLat; lat += latitudeStep {
		for lng := minLng; lng <= maxLng; lng += longitudeStep {
			if g.isPointInPolygon(lat, lng) {
				points = append(points, []float64{lat, lng})
			}
		}
	}
	return points
}

// بررسی اینکه آیا نقطه‌ای درون پلی‌گون قرار دارد یا خیر
func (g *GeoHashGenerator) isPointInPolygon(lat, lng float64) bool {
	inside := false
	numPoints := len(g.polygon)
	for i, j := 0, numPoints-1; i < numPoints; j = i {
		xi, yi := g.polygon[i][0], g.polygon[i][1]
		xj, yj := g.polygon[j][0], g.polygon[j][1]

		// محاسبه تقاطع نقطه با پلی‌گون
		intersect := (yi > lng) != (yj > lng) && lat < (xj-xi)*(lng-yi)/(yj-yi)+xi
		if intersect {
			inside = !inside
		}
	}
	return inside
}

// کدگذاری مختصات جغرافیایی به GeoHash
func (g *GeoHashGenerator) encodeGeoHash(lat, lng float64, precision int) string {
	minLat, maxLat := -90.0, 90.0
	minLng, maxLng := -180.0, 180.0
	var geoHash string
	even := true
	bit := 0
	ch := 0

	// حلقه برای کدگذاری مختصات به GeoHash
	for len(geoHash) < precision {
		if even {
			mid := (minLng + maxLng) / 2
			if lng > mid {
				ch |= 1 << (4 - bit)
				minLng = mid
			} else {
				maxLng = mid
			}
		} else {
			mid := (minLat + maxLat) / 2
			if lat > mid {
				ch |= 1 << (4 - bit)
				minLat = mid
			} else {
				maxLat = mid
			}
		}
		even = !even
		if bit < 4 {
			bit++
		} else {
			geoHash += string(geoHashChars[ch])
			bit = 0
			ch = 0
		}
	}

	return geoHash
}
