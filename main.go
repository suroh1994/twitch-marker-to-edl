package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
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

	csvFilename := path.Base(csvFile)[:len(path.Ext(csvFile))*-1]
	edlFilename := fmt.Sprintf("%s%s.edl", path.Dir(csvFile), csvFilename)
	edl.ExportToFile(edlFilename)
}

func convertExportToEdl(export TwitchStreamMarkerExport) (edl EDLFormat) {
	edl.Title = "Twitch Stream Markers"
	edl.FCM = FCM_NON_DROP_FRAME
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
