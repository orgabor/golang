package parts

import "fmt"

type Monitor struct {
	Resolution string
	Connector  string
	Value      float64
}

func (m Monitor) Specs() string {
	return fmt.Sprintf("Monitor\nResolution: %s\nConnector: %s", m.Resolution, m.Connector)
}

func (m Monitor) Price() string {
	return fmt.Sprintf("$%0.2f", m.Value)
}

type HardDrive struct {
	Type      string
	Connector string
	Value     float64
}

func (h HardDrive) Specs() string {
	return fmt.Sprintf("Hard Drive\nType: %s\nConnector: %s", h.Type, h.Connector)
}

func (h HardDrive) Price() string {
	return fmt.Sprintf("$%0.2f", h.Value)
}

type Keyboard struct {
	Keys     int
	Switches string
	Value    float64
}

func (k Keyboard) Specs() string {
	return fmt.Sprintf("Keyboard\nKeys: %d\nSwitch Type: %s", k.Keys, k.Switches)
}

func (k Keyboard) Price() string {
	return fmt.Sprintf("$%0.2f", k.Value)
}
