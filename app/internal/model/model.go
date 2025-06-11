package model

import (
	"sync"
	"time"

	"github.com/gdamore/tcell/v2"
)

type GameState struct {
	World        World
	Symptoms     []Symptom
	Gameticker   *time.Ticker
	CurrentScene Scene
	Mutex        sync.Mutex
}

type Scene interface {
	Update()
	Draw(s tcell.Screen)
	HandleEvent(ev tcell.Event)
}

type World struct {
	Total    int
	Healthy  int
	Infected int
	Dead     int

	NewInfected int
	NewDead     int

	Disease    Disease
	Regions    []Region
	DaysPassed int

	Credit int

	Speed time.Duration
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
}

type Symptom struct {
	ID                int
	Name              string
	MortalityBonus    float64
	TransmissionBonus float64
	Cost              int
	Unlocked          bool
}
