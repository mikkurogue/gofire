package main

import (
	"fmt"
	"gofire/tracker"
	"gofire/ui"
	"log"
	"time"
)

func main() {

	fmt.Println("hello world")

	tracker := tracker.InitTracker()
	err := tracker.LoadData("gametimes.json")
	if err != nil {
		log.Println("no exisiting data found, starting a new tracker!")
	}

	// somewhere we need to compile a list of executables for time tracking
	knownGames := map[string]bool{
		"WoW.exe": true,
	}

	fmt.Println("Creating window")
	// Create and show the UI window
	window := ui.CreateWindow("GoFire", 1280, 720)
	go func() {

		GetInitialState(window, tracker, knownGames)

		// NOTE: Make sure we find the correct timer for the loop updates.
		for range time.Tick(1 * time.Second) {
			tracker.UpdateGameTimes(knownGames)
			tracker.SaveData("gametimes.json")

			// Update the label with the current game times
			gameName := tracker.RunningGameName()

			fmt.Println(gameName)

			window.UpdateLabel(gameName)
		}
	}()
	window.Show()
}

func GetInitialState(window *ui.Window, tracker *tracker.GameTimeTracker, knownGames map[string]bool) {
	tracker.UpdateGameTimes(knownGames)

	tracker.SaveData("gametimes.json")

	gameName := tracker.RunningGameName()
	window.UpdateLabel(gameName)

}
