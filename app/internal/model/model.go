package model

import "github.com/gdamore/tcell/v2"

type Scene interface {
	Update()
	Draw(s tcell.Screen)
	EventHandler(ev tcell.Event)
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
