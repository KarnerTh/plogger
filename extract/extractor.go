package extract

import "fmt"

type Data struct {
	LogLine string
	Value   float64
}

type Extractor interface {
	Key() string
	Extract(logLine string) (Data, error)
}

var availableExtractors = []Extractor{
	NewPingExtrator(),
	NewJestHeapExtrator(),
}

func GetExtractor(key string) (Extractor, error) {
	for _, item := range availableExtractors {
		if item.Key() == key {
			return item, nil
		}
	}

	return nil, fmt.Errorf("No extractor available with key '%s'", key)
}
