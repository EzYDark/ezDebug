package main

import (
	"flag"

	"github.com/ezydark/ezDebugTUI/src/tui"
	"github.com/rs/zerolog/log"
)

const (
	WindowWidth  = 450
	WindowHeight = 900
	WindowTitle  = "ezChat"
)

func main() {
	exampleFlag := flag.Bool("example", false, "Run in debug TUI mode")
	flag.Parse()

	if *exampleFlag {
		app := tui.InitDebugTUI()
		tui.InitWithWriter(app.LogWriter())
		log.Info().Msgf("Starting %s in debug TUI mode...", WindowTitle)
		if err := app.Run(); err != nil {
			log.Fatal().Msgf("Error running debug TUI:\n%v", err)
		}
	}
}
