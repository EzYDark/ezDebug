// Start this example application with the '-example' flag
// to see how ezDebugTUI works.

package main

import (
	"flag"

	"github.com/ezydark/ezDebugTUI/tui"
	"github.com/ezydark/ezLog"
	"github.com/ezydark/ezLog/log"
)

func main() {
	exampleFlag := flag.Bool("example", false, "Show example of ezDebugTUI usage")
	flag.Parse()

	// Initialize the global logger with custom
	// zerolog configuration
	ezLog.Init()

	if *exampleFlag {
		debugTUI := tui.Init()

		// Have to reinitialize the global logger with the TUI's Writer
		// otherwise the UI glitches when redrawing
		ezLog.InitWithWriter(debugTUI.GetLogWriter())

		log.Info().Msg("Starting example of ezDebugTUI usage...")

		features := []tui.Feature{
			{
				Name:        "Feature 1",
				Description: "Description of feature 1",
				Enabled:     false,
				Action: func(enabled bool) {
					log.Info().Msgf("Feature 1 is %v", enabled)
				},
			},
			{
				Name:        "Feature 2",
				Description: "Description of feature 2",
				Enabled:     false,
				Action: func(enabled bool) {
					log.Info().Msgf("Feature 2 is %v", enabled)
				},
			},
		}
		tui.GetFeatureList().Set(features)

		// Start DebugTUI application (Blocking call)
		if err := debugTUI.Start(); err != nil {
			log.Fatal().Msgf("Error running ezDebugTUI example:\n%v", err)
		}
	} else {
		log.Fatal().Msg("Run with '-example' flag to start the example usage of ezDebugTUI.")
	}
}
