package main

import (
	"flag"

	"github.com/ezydark/ezDebugTUI/src"
	"github.com/rs/zerolog/log"
)

func main() {
	exampleFlag := flag.Bool("example", false, "Show example of ezDebugTUI usage")
	flag.Parse()

	if *exampleFlag {
		app := src.InitDebugTUI()
		// InitLogger(app.GetLogWriter())

		log.Info().Msg("Starting example of ezDebugTUI usage...")

		features := []src.Feature{
			{
				Name:        "Feature 111",
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
		src.GetFeatureList().Set(features)

		// Start TUI application (Blocking call)
		if err := app.Start(); err != nil {
			log.Fatal().Msgf("Error running ezDebugTUI example:\n%v", err)
		}
	} else {
		log.Fatal().Msg("Run with '-example' flag to start the example usage of ezDebugTUI.")
	}
}
