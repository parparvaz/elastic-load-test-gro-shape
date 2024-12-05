# Elastic Load Test GEO Shape

# GeoShapes Project - Commands Overview

This project includes four commands that help in generating, inserting, and testing geospatial data with Elasticsearch. Below is an overview of each command:


## GeoShapes Project - Command 1: `geo-shape make-index`

### Overview
The `geo-shape make-index` command creates an index in Elasticsearch with the name `geo-shapes-v1`. This index is used for storing geospatial data with a specific mapping configuration.

### Mapping Configuration
The index is created with the following mappings:

```json
{
  "mappings": {
    "geo-shapes-v1-mappings": {
      "properties": {
        "location": {
          "type": "geo_shape"
        },
        "shop_id": {
          "type": "integer"
        },
        "polygon_id": {
          "type": "integer"
        },
        "radius_base": {
          "type": "boolean"
        }
      }
    }
  }
}
```

#### Description of Fields
- **location**: A `geo_shape` field used to store geospatial data (e.g., polygons, circles, etc.).
- **shop_id**: An integer representing the unique identifier for a shop.
- **polygon_id**: An integer representing the unique identifier for a polygon associated with a shop.
- **radius_base**: A boolean field indicating whether the geo-shape is based on a radius or not.

### Usage
To run this command, use the following command in your terminal:

```bash
go run main.go geo-shape make-index
```

This will create the `geo-shapes-v1` index in your Elasticsearch instance with the specified mappings.

### Requirements
- Elasticsearch instance running

## GeoShapes Project - Command 2: `geo-shape make-fake-polygons`

### Overview
The `geo-shape make-fake-polygons` command generates 10,000 random polygons within the city of Tehran. These polygons are saved into a file named `fake_polygons.json`.

### How It Works
- The command generates random polygons within the city limits of Tehran.
- A total of 10,000 polygons are generated.
- The generated polygons are stored in the file `fake_polygons.json`.

### Usage
To run this command, use the following command in your terminal:

```bash
go run main.go geo-shape make-fake-polygons
```
### Output
The command will generate a file named `fake_polygons.json` in the root directory. This file contains 10,000 randomly generated polygons in JSON format.

#### Example of the JSON format:
```json
[
  {
    "polygon_id": 1,
    "coordinates": [
      [51.378, 35.694],
      [51.380, 35.695],
      [51.381, 35.693],
      [51.379, 35.692],
      [51.378, 35.694]
    ],
    "shop_id": 101,
    "radius_base": true
  },
  {
    "polygon_id": 2,
    "coordinates": [
      [51.389, 35.700],
      [51.390, 35.701],
      [51.391, 35.699],
      [51.389, 35.698],
      [51.389, 35.700]
    ],
    "shop_id": 102,
    "radius_base": false
  }
]
```
- **polygon_id**: A unique identifier for each generated polygon.
- **coordinates**: The list of geographical coordinates (latitude, longitude) that define the polygon shape. The coordinates are generated randomly within Tehran’s geographical boundaries.
- **shop_id**: A randomly assigned shop ID associated with the polygon.
- **radius_base**: A boolean indicating whether the polygon is radius-based.

## GeoShapes Project - Command 3: `geo-shape insert-fake-polygon-to-elastic`

### Overview
The `geo-shape insert-fake-polygon-to-elastic` command reads the `fake_polygons.json` file, randomly selects values for `polygon_id`, `shop_id`, and `is_radius_base`, and inserts the data into the `geo-shaps-v1` index in Elasticsearch using the Bulk API.

### How It Works
- The command reads the `fake_polygons.json` file.
- Randomly selects values for `polygon_id`, `shop_id`, and `is_radius_base`.
- Inserts the selected data into Elasticsearch using the Bulk API for efficient insertion.

### Usage
To run this command, use the following command in your terminal:

```bash
go run main.go geo-shape insert-fake-polygon-to-elastic
```

### Input
The input for this command is the `fake_polygons.json` file, which contains the randomly generated polygons.

#### Example of the JSON format:

```json
[
  {
    "polygon_id": 1,
    "coordinates": [
      [51.378, 35.694],
      [51.380, 35.695],
      [51.381, 35.693],
      [51.379, 35.692],
      [51.378, 35.694]
    ],
    "shop_id": 101,
    "radius_base": true
  },
  {
    "polygon_id": 2,
    "coordinates": [
      [51.389, 35.700],
      [51.390, 35.701],
      [51.391, 35.699],
      [51.389, 35.698],
      [51.389, 35.700]
    ],
    "shop_id": 102,
    "radius_base": false
  }
]
```
### Output
The command will insert the data from the `fake_polygons.json` file into the `geo-shaps-v1` index in Elasticsearch.  
Data is inserted in bulk for improved performance.


## GeoShapes Project - Command 4: `geo-shape load-test`

### Overview
The `geo-shape load-test` command is designed to test the load capacity of Elasticsearch by sending a configurable number of requests in rapid succession. It monitors system performance, including CPU and RAM usage, during the test, and saves the results in three separate Excel files for further analysis.

### How It Works
- The command sends a specified number of requests to Elasticsearch per millisecond.
- Simultaneously, it monitors the system's CPU, RAM, and other performance metrics.
- The performance data is saved into three separate Excel files for analysis:
    - `Elastic Load Test.xlsx`
    - `Elastic Monitoring.xlsx`
    - `Local Monitoring.xlsx`

### Usage
To run this command, use the following command in your terminal:

```bash
go run main.go geo-shape load-test
```
### Configuration
The number of requests to be sent to Elasticsearch can be configured via the controller `geo-shape.controller.go`.  
The test duration and other parameters can be adjusted based on system and Elasticsearch settings.

### Output
The following three Excel files are generated as output:

1. `Elastic Load Test.xlsx`: Contains data about the load test, including the number of requests and responses.
2. `Elastic Monitoring.xlsx`: Provides system and Elasticsearch performance metrics during the test.
3. `Local Monitoring.xlsx`: Contains local system performance metrics (CPU, RAM usage, etc.) during the test.
