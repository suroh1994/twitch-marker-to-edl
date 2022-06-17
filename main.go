package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var version = "0.0.0"

func main() {
	fmt.Printf("Twitch2Davinci version %s\n", version)

	if len(os.Args) < 2 {
		fmt.Printf("incorrect usage: program csvfile")
		return
	}

	csvFile := os.Args[1]
	export, err := ImportCSV(csvFile)
	if err != nil {
		fmt.Printf("failed to import Twitch stream marker export: %v\n", err)
		return
	}

	edl := convertExportToEdl(export)

	edlFile := generateOutputPath(csvFile)
	edl.ExportToFile(edlFile)
}

func convertExportToEdl(export TwitchStreamMarkerExport) (edl EDLFormat) {
	edl.Title = "Twitch Stream Markers"
	edl.FCM = fcmNonDropFrame
	for _, marker := range export.Markers {
		edl.TimelineMarkers = append(edl.TimelineMarkers, TimelineMarker{
			Title:     marker.Title,
			Timestamp: marker.Timestamp,
			Color:     Blue,
			Duration:  1,
		})
	}
	return
}

func generateOutputPath(inputPath string) string {
	filenameWithExtension := filepath.Base(inputPath)
	fileExtension := filepath.Ext(inputPath)

	filename := filenameWithExtension[:len(filenameWithExtension)-len(fileExtension)]
	return filepath.Join(filepath.Dir(inputPath), filename+".edl")
}
