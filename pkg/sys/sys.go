package sys

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/wailsapp/wails"
)

type Stats struct {
	log *wails.CustomLogger
}

type CPUUsage struct {
	Average int `json:"avg"`
}

func (s *Stats) WailsInit(runtime *wails.Runtime) error {
	s.log = runtime.Log.New("Stats")

	go func() {
		for {
			runtime.Events.Emit("cpu_usage", s.GetCPUUsage())
			time.Sleep(time.Second)
		}
	}()
	return nil
}

func (s *Stats) GetCPUUsage() *CPUUsage {
	percents, err := cpu.Percent(time.Second, false)
	if err != nil {
		s.log.Errorf("unable to get cpu stats: %s", err.Error())
	}

	return &CPUUsage{
		Average: int(percents[0]),
	}
}
