package tracker

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/shirou/gopsutil/process"
)

type GameTimeTracker struct {
	GameTimes  map[string]time.Duration `json:"game_times"`
	StartTimes map[string]time.Time     `json: "-"`
}

func InitTracker() *GameTimeTracker {
	return &GameTimeTracker{
		GameTimes:  make(map[string]time.Duration),
		StartTimes: make(map[string]time.Time),
	}
}

func (gtt *GameTimeTracker) SaveData(filename string) error {

	// convert duration to minutes for json output
	gameTimesInMinutes := make(map[string]int)
	for game, duration := range gtt.GameTimes {
		gameTimesInMinutes[game] = int(duration.Minutes())
	}

	data, err := json.MarshalIndent(map[string]interface{}{
		"game_times": gameTimesInMinutes,
	}, "", " ")
	if err != nil {
		return fmt.Errorf("error marshaling data: %v", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	return nil
}

func (gtt *GameTimeTracker) LoadData(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	// temporarily load json data
	var temp struct {
		GameTimes map[string]int `json: "game_times"`
	}

	err = json.Unmarshal(data, &temp)
	if err != nil {
		return fmt.Errorf("error unmarhsaling data: %v", err)
	}

	for game, minutes := range temp.GameTimes {
		gtt.GameTimes[game] = time.Duration(minutes) * time.Minute
	}

	return nil
}

func (gtt *GameTimeTracker) UpdateGameTimes(knownGames map[string]bool) {

	processes, err := process.Processes()
	if err != nil {
		log.Fatalf("error getting processes: %v", err)
	}

	runningGames := make(map[string]bool)
	for _, prc := range processes {
		name, err := prc.Name()
		if err != nil {
			continue
		}

		if knownGames[name] {
			runningGames[name] = true

			if _, exists := gtt.StartTimes[name]; !exists {
				startTime, err := prc.CreateTime()
				if err != nil {
					continue
				}

				unixTime := time.Unix(startTime/1000, 0)
				gtt.StartTimes[name] = unixTime
			} else {
				gtt.GameTimes[name] += time.Minute
			}
		}
	}
}

func (gtt *GameTimeTracker) RunningGameName() string {
	var result string

	for game := range gtt.GameTimes {
		result = fmt.Sprintf("Playing: %s\n", game)
	}

	return result
}
