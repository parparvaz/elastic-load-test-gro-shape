package domain

import "time"

type ScenarioLoadTest struct {
	WholeTime    int `yaml:"wholeTime"`
	Interval     int `yaml:"interval"`
	RequestCount int `yaml:"requestCount"`
}

type ScenarioLocalMonitoring struct {
	ScenarioName string
	Data         []LocalMonitoring
	Start        time.Time
	End          time.Time
}

type ScenarioElasticMonitoring struct {
	ScenarioName string
	Data         []ElasticMonitoring
	Start        time.Time
	End          time.Time
}
