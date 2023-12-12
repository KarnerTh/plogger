package extract

import (
	"log/slog"
	"regexp"
	"strconv"
)

type jestHeapExtractor struct{}

var regHeap = regexp.MustCompile(`\((?P<heap>.*)\sMB.*\)`)

func NewJestHeapExtrator() Extractor {
	return jestHeapExtractor{}
}

func (e jestHeapExtractor) Key() string {
	return "jest_heap"
}

func (e jestHeapExtractor) Extract(logLine string) (Data, error) {
	groups := regHeap.FindStringSubmatch(logLine)
	resultIndex := regHeap.SubexpIndex("heap")

	if len(groups) == 0 {
		// TODO: empty values should not be written
		return Data{LogLine: logLine, Value: 0}, nil
	}

	heapResult := groups[resultIndex]
	heapSize, err := strconv.ParseFloat(heapResult, 64)
	if err != nil {
		slog.Error("Error in parsing heap size result", slog.Any("error", err))
		// TODO: empty values should not be written
		return Data{LogLine: logLine, Value: 0}, err
	}

	return Data{LogLine: logLine, Value: heapSize}, nil
}
