package logic

import (
	"math"
	"worldinc/app/internal/model"
)

func DoWorldTick(w *model.World) {
	w.DaysPassed++
	w.NewInfected = 0
	w.NewDead = 0
	w.Total = w.Healthy + w.Infected + w.Dead // it should be const btw

	// i dunno if i should just pass values directly, cuz those are kinda helpful when it comes to debug the shit
	th := int(math.Ceil(w.Disease.Transmission * float64(w.Healthy))) // th stand for Transmission * Healthy
	mi := int(math.Ceil(w.Disease.Mortality * float64(w.Infected)))   // mi stand for Transmission * Healthy

	if w.Infected == 0 {
		th = 0
	}

	w.NewInfected = th
	w.Infected += w.NewInfected
	w.Healthy -= w.NewInfected
	w.Credit += int(math.Ceil(float64(w.NewInfected) / 1000))

	if w.Healthy <= 0 {
		w.NewInfected = w.Healthy + w.NewInfected
		w.Healthy = 0
	}

	if w.DaysPassed > 5 { // HARDCODED?? TODO: Pass an estimate value of safe days to prevent mass dying from the start
		w.NewDead = mi
		w.Dead += w.NewDead
		w.Infected -= w.NewDead

		if w.Infected <= 0 {
			w.NewDead = w.Infected + w.NewDead
			w.Infected = 0
		}
	}

	// +1 credt every 5 days
	if w.DaysPassed%5 == 0 {
		w.Credit += 1
	}

	// fmt.Printf("%v. H: %v I: %v [%v] D: %v [%v]  ", w.DaysPassed, w.Healthy, w.Infected, w.NewInfected, w.Dead, w.NewDead)
	// fmt.Printf("T: %v M: %v T*H: %v M*I: %v\n", w.Disease.Transmission, w.Disease.Mortality, th, mi)
	// time.Sleep(time.Second * 5 / 20)
}
