package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"log"
)

type TwitchStreamMarkerExport struct {
	Markers []TwitchStreamMarker
}

type TwitchStreamMarker struct {
	Timestamp   time.Time
	CreatorRole string
	CreatorName string
	Title       string
}

const (
	expectedColumnCount = 4
	csvTimestampFormat  = "15:04:05"
)

type ColumnCountError int

func (e ColumnCountError) Error() string {
	return fmt.Sprintf("input issue: expected %d columns but got %d", expectedColumnCount, int(e))
}

func ImportCSV(filePath string) (export TwitchStreamMarkerExport, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("failed to open file %q: %v\n", filePath, err)
		return
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		fmt.Printf("failed to read csv file %q: %v\n", filePath, err)
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
	if len(line) != expectedColumnCount {
		err = ColumnCountError(len(line))
		return
	}

	marker.Timestamp, err = time.Parse(csvTimestampFormat, line[0])
	if err != nil {
		return
	}

	marker.CreatorRole = line[1]
	marker.CreatorName = line[2]
	marker.Title = line[3]
	return
}
