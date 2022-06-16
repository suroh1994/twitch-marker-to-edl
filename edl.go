package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type EDLFormat struct {
	Title           string
	FCM             string
	TimelineMarkers []TimelineMarker
}

const (
	fcmNonDropFrame    = "NON-DROP FRAME"
	titleLineFormat    = "TITLE: %s\n"
	fcmLineFormat      = "FCM: %s\n"
	edlTimestampFormat = "15:04:05:00"
	markerLine1Format  = "%03d  001      V     C        %s %s %s %s  \n"
	markerLine2Format  = " |C:%s |M:%s |D:%d"
)

type TimelineMarker struct {
	Timestamp time.Time
	Color     MarkerColor
	Title     string
	Duration  int
}

type MarkerColor string

const (
	Blue     MarkerColor = "ResolveColorBlue"
	Cyan     MarkerColor = "ResolveColorCyan"
	Green    MarkerColor = "ResolveColorGreen"
	Yellow   MarkerColor = "ResolveColorYellow"
	Red      MarkerColor = "ResolveColorRed"
	Pink     MarkerColor = "ResolveColorPink"
	Purple   MarkerColor = "ResolveColorPurple"
	Fuchsia  MarkerColor = "ResolveColorFuchsia"
	Rose     MarkerColor = "ResolveColorRose"
	Lavender MarkerColor = "ResolveColorLavender"
	Sky      MarkerColor = "ResolveColorSky"
	Mint     MarkerColor = "ResolveColorMint"
	Lemon    MarkerColor = "ResolveColorLemon"
	Sand     MarkerColor = "ResolveColorSand"
	Cocoa    MarkerColor = "ResolveColorCocoa"
	Cream    MarkerColor = "ResolveColorCream"
)

func (edl EDLFormat) ExportToFile(filepath string) {
	f, err := os.Create(filepath)
	if err != nil {
		fmt.Printf("failed to create file %q: %v\n", filepath, err)
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	var output []string

	// print title line
	output = append(output, fmt.Sprintf(titleLineFormat, edl.Title))

	// print FCM line
	output = append(output, fmt.Sprintf(fcmLineFormat, edl.FCM))

	// for each marker:
	for idx, marker := range edl.TimelineMarkers {
		// print empty line
		output = append(output, "\n")

		// print index, 001, V, C, start time, end time, start time, end time
		startTime := marker.Timestamp
		endTime := marker.Timestamp.Add(time.Duration(marker.Duration) * time.Second)
		output = append(output, fmt.Sprintf(markerLine1Format, idx+1, startTime.Format(edlTimestampFormat), endTime.Format(edlTimestampFormat), startTime.Format(edlTimestampFormat), endTime.Format(edlTimestampFormat)))

		// print |C:color |M: title |D:duration
		output = append(output, fmt.Sprintf(markerLine2Format, marker.Color, marker.Title, marker.Duration))
	}

	for _, line := range output {
		_, err = w.WriteString(line)
		if err != nil {
			fmt.Printf("failed to write line %q : %v\n", line, err)
			return
		}
	}
}
