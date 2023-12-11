package extract

import (
	"log/slog"
	"regexp"
	"strconv"
)

type pingExtractor struct{}

var regTime = regexp.MustCompile(`time=(?P<time>.*)\sms`)

func NewPingExtrator() Extractor {
	return pingExtractor{}
}

func (e pingExtractor) Key() string {
	return "ping"
}

func (e pingExtractor) Extract(logLine string) (Data, error) {
	groups := regTime.FindStringSubmatch(logLine)
	resultIndex := regTime.SubexpIndex("time")

	if len(groups) == 0 {
		// TODO: empty values should not be written
		return Data{LogLine: logLine, Value: 0}, nil
	}

	timeResult := groups[resultIndex]
	time, err := strconv.ParseFloat(timeResult, 64)
	if err != nil {
		slog.Error("Error in parsing time result", slog.Any("error", err))
		// TODO: empty values should not be written
		return Data{LogLine: logLine, Value: 0}, err
	}

	return Data{LogLine: logLine, Value: time}, nil
}
