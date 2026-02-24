package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Charger represents a single charger at a station.
type Charger struct {
	ChargerID string `json:"charger_id"`
	Type      string `json:"type"`
}

// Station represents a charging station (site) with one or more chargers.
type Station struct {
	Name     string    `json:"name"`
	Address1 string    `json:"address1"`
	City     string    `json:"city"`
	State    string    `json:"state"`
	Network  string    `json:"network"`
	Chargers []Charger `json:"chargers"`
}

func stationKey(s Station) string {
	return strings.Join([]string{s.Address1, s.City, s.State, s.Network}, "|")
}

func SecondaryCoalescedName(s Station) string {
	parts := strings.Split(s.Name, ", ")
	return strings.TrimSpace(parts[1])
}

func GroupStations(stations []Station) []Station {
	if len(stations) == 0 {
		return nil
	}

	byKey := make(map[string]*Station)

	for i := range stations {
		s := &stations[i]
		key := stationKey(*s)

		if existing, ok := byKey[key]; ok {
			existing.Chargers = append(existing.Chargers, s.Chargers...)
			existing.Name = existing.Name + ", " + s.Name
		} else {
			merged := Station{
				Name:     s.Name,
				Address1: s.Address1,
				City:     s.City,
				State:    s.State,
				Network:  s.Network,
				Chargers: append([]Charger(nil), s.Chargers...),
			}
			byKey[key] = &merged
		}
	}

	out := make([]Station, 0, len(byKey))
	for _, st := range byKey {
		out = append(out, *st)
	}
	return out
}

func main() {
	input := []Station{
		{
			Name:     "Springfield 3805 Moto-1",
			Address1: "3805 Moto St",
			City:     "Springfield",
			State:    "IL",
			Network:  "Chargestop",
			Chargers: []Charger{{ChargerID: "3805 Moto-1", Type: "CCS"}},
		},
		{
			Name:     "Springfield 3805 Moto-2",
			Address1: "3805 Moto St",
			City:     "Springfield",
			State:    "IL",
			Network:  "Chargestop",
			Chargers: []Charger{{ChargerID: "3805 Moto-2", Type: "CCS"}},
		},
		{
			Name:     "Springfield 111 Main",
			Address1: "111 Main St",
			City:     "Springfield",
			State:    "IL",
			Network:  "Chargego",
			Chargers: []Charger{
				{ChargerID: "Hall", Type: "CCS"},
				{ChargerID: "Oates", Type: "CCS"},
			},
		},
		{
			Name:     "Springfield 555 Broadway-1",
			Address1: "555 Broadway",
			City:     "Springfield",
			State:    "IL",
			Network:  "Chargestop",
			Chargers: []Charger{{ChargerID: "Broadway-1", Type: "CCS"}},
		},
		{
			Name:     "Springfield 555 Broadway-2",
			Address1: "555 Broadway",
			City:     "Springfield",
			State:    "IL",
			Network:  "Chargestop",
			Chargers: []Charger{{ChargerID: "Broadway-2", Type: "CCS"}},
		},
		{
			Name:     "Springfield 3805 Moto-3",
			Address1: "3805 Moto St",
			City:     "Springfield",
			State:    "IL",
			Network:  "Chargestop",
			Chargers: []Charger{{ChargerID: "3805 Moto-3", Type: "CCS"}},
		},
	}

	grouped := GroupStations(input)

	for _, st := range grouped {
		_ = SecondaryCoalescedName(st)
	}

	jsonOut, _ := json.MarshalIndent(grouped, "", "    ")
	fmt.Println("Grouped stations (desired structure):")
	fmt.Println(string(jsonOut))
}
