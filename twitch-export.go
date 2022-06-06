package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"

	"log"
)

type TwitchStreamMarkerExport struct {
	Markers []TwitchStreamMarker
}

type TwitchStreamMarker struct {
	Timestamp   time.Duration
	CreatorRole string
	CreatorName string
	Title       string
}

const (
	ExpectedColumnCount           = 4
	ExpectedTimestampSegmentCount = 3
)

type ColumnCountError int

func (e ColumnCountError) Error() string {
	return fmt.Sprintf("input issue: expected %d columns but got %d", ExpectedColumnCount, int(e))
}

type TimestampSegmentCountError int

func (e TimestampSegmentCountError) Error() string {
	return fmt.Sprintf("input issue: expected %d timestamp segments but got %d", ExpectedTimestampSegmentCount, int(e))
}

func ImportCSV(filePath string) (export TwitchStreamMarkerExport, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		return
	}

	return parseCSV(data)
}

func parseCSV(content [][]string) (export TwitchStreamMarkerExport, err error) {
	for idx, line := range content {
		marker, err := parseLine(line)
		if err != nil {
			log.Printf("failed to process line %d: %v", idx, err)
		} else {
			export.Markers = append(export.Markers, marker)
		}
	}
	return
}

func parseLine(line []string) (marker TwitchStreamMarker, err error) {
	if len(line) != ExpectedColumnCount {
		err = ColumnCountError(len(line))
		return
	}

	timestampSegments := strings.Split(line[0], ":")
	if len(timestampSegments) != ExpectedTimestampSegmentCount {
		err = TimestampSegmentCountError(len(timestampSegments))
		return
	}

	adjustedTimeString := fmt.Sprintf("%sh%sm%ss", timestampSegments[0], timestampSegments[1], timestampSegments[2])
	marker.Timestamp, err = time.ParseDuration(adjustedTimeString)
	if err != nil {
		return
	}

	marker.CreatorRole = line[1]
	marker.CreatorName = line[2]
	marker.Title = line[3]
	return
}
