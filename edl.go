package main

import "time"

type EDLFormat struct {
	Title           string
	FCM             string
	TimelineMarkers []TimelineMarker
}

type TimelineMarker struct {
	Timestamp time.Duration
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
