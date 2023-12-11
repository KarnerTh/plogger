package extract

type Data struct {
	LogLine string
	Value   float64
}
type Extractor interface {
	Key() string
	Extract(logLine string) (Data, error)
}
