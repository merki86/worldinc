package model

import (
	"sync"

	"github.com/gdamore/tcell/v2"
)

type GameState struct {
	World        World
	CurrentScene Scene
	Mutex        sync.Mutex
}

type Scene interface {
	Update(done chan struct{})
	Draw(s tcell.Screen, done chan struct{})
	HandleEvent(s tcell.Screen)
	Next() Scene
}

type World struct {
	Population int
	Infected   int
	Dead       int
	Regions    []Region
	Disease    Disease
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
	Name           string
	MortalityBonus float64
	Cost           int
	Unlocked       bool
}
