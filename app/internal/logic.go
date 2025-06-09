package internal

import (
	"math"
	"worldinc/app/internal/model"
)

func Simulate(world *model.World) (int, int) {
	world.DaysPassed++
	infectedToday := 0
	diedToday := 0

	//total := world.Healthy + world.Infected + world.Dead // it is const de-facto

	th := math.Ceil(world.Disease.Transmission * float64(world.Healthy))
	mi := math.Ceil(world.Disease.Mortality * float64(world.Infected))

	infectedToday = int(th)
	world.Infected += infectedToday
	world.Healthy -= infectedToday
	if world.Healthy <= 0 {
		infectedToday = world.Healthy + infectedToday
		world.Healthy = 0
	}

	if world.DaysPassed > 5 {
		diedToday = int(mi)
		world.Dead += diedToday
		world.Infected -= diedToday
		if world.Infected <= 0 {
			diedToday = world.Infected + diedToday
			world.Infected = 0
		}
	}

	return infectedToday, diedToday

	// fmt.Printf("%v. H: %v I: %v [%v] D: %v [%v]  ", world.DaysPassed, world.Healthy, world.Infected, infectedToday, world.Dead, diedToday)
	// fmt.Printf("T: %v M: %v T*H: %v M*I: %v\n", world.Disease.Transmission, world.Disease.Mortality, th, mi)
	// time.Sleep(time.Second * 5 / 20)
}
