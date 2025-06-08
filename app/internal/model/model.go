package model

import (
	"sync"
	"time"

	"github.com/gdamore/tcell/v2"
)

type GameState struct {
	World        World
	CurrentScene Scene
	Mutex        sync.Mutex
}

type Scene interface {
	Update(dt time.Duration)
	Draw(s tcell.Screen)
	HandleEvent(ev tcell.Event)
	// Next() Scene
}

type World struct {
	Population int
	Infected   int
	Dead       int
	Disease    Disease
	Regions    []Region
	DaysPassed int
}

type Region struct {
	Name       string
	Population int
	Infected   int
	Dead       int
}

type Disease struct {
	Name         string
	Mortality    float64
	Transmission float64
	Discovered   bool
	Symptoms     []Symptom
}

type Symptom struct {
	Name              string
	MortalityBonus    float64
	TransmissionBonus float64
	Cost              int
	Unlocked          bool
}
