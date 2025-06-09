package internal

import (
	"math"
	"worldinc/app/internal/model"
)

func Simulate(world *model.World) {
	world.DaysPassed++
	world.NewInfected = 0
	world.NewDead = 0

	world.Total = world.Healthy + world.Infected + world.Dead // it should be const btw

	th := math.Ceil(world.Disease.Transmission * float64(world.Healthy))
	mi := math.Ceil(world.Disease.Mortality * float64(world.Infected))

	world.NewInfected = int(th)
	world.Infected += world.NewInfected
	world.Healthy -= world.NewInfected
	if world.Healthy <= 0 {
		world.NewInfected = world.Healthy + world.NewInfected
		world.Healthy = 0
	}

	if world.DaysPassed > 5 {
		world.NewDead = int(mi)
		world.Dead += world.NewDead
		world.Infected -= world.NewDead
		if world.Infected <= 0 {
			world.NewDead = world.Infected + world.NewDead
			world.Infected = 0
		}
	}

	// fmt.Printf("%v. H: %v I: %v [%v] D: %v [%v]  ", world.DaysPassed, world.Healthy, world.Infected, world.NewInfected, world.Dead, world.NewDead)
	// fmt.Printf("T: %v M: %v T*H: %v M*I: %v\n", world.Disease.Transmission, world.Disease.Mortality, th, mi)
	// time.Sleep(time.Second * 5 / 20)
}
