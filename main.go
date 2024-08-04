package main

import (
	"fmt"
	"gofire/tracker"
	"log"
	"time"

	"github.com/shirou/gopsutil/process"
)

func main() {

	tracker := tracker.InitTracker()
	err := tracker.LoadData("gametimes.json")
	if err != nil {
		log.Println("no exisiting data found, starting a new tracker!")
	}

	// somewhere we need to compile a list of executables for time tracking
	knownGames := map[string]bool{
		"WoW.exe": true,
	}

	for {
		processes, err := process.Processes()
		if err != nil {
			log.Fatalf("Error getting processes: %v", err)
		}

		// track the running processes
		runningGames := make(map[string]bool)

		for _, proc := range processes {
			name, err := proc.Name()
			if err != nil {
				continue
			}

			if knownGames[name] {

				runningGames[name] = true

				fmt.Print("detected game ", name)

				if _, exists := tracker.StartTimes[name]; !exists {
					// if game process did not first exist, then give it a start time
					fmt.Println("process no recorded times, creating new entry...")
					startTime, err := proc.CreateTime()
					if err != nil {
						continue
					}

					// Convert milliseconds to time.Time
					unixStartTime := time.Unix(startTime/1000, 0)
					tracker.StartTimes[name] = unixStartTime
				} else {
					tracker.GameTimes[name] += time.Minute
				}
			}
		}

		// remove games that have been closed down
		for game := range tracker.StartTimes {
			if !runningGames[game] {
				delete(tracker.StartTimes, game)
			}
		}

		// Save data every minute
		time.Sleep(1 * time.Minute)
		err = tracker.SaveData("gametimes.json")
		if err != nil {
			log.Fatalf("Error saving data: %v", err)
		}
	}
}
