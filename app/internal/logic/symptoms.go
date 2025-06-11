package logic

import (
	"errors"
	"worldinc/app/internal/model"
)

func BuySymptom(id int, g *model.GameState) (*model.Symptom, error) {
	for i, v := range g.Symptoms {
		if v.ID == id {
			if g.World.Credit-v.Cost < 0 {
				return nil, errors.New("not enough credits") // TODO: no credit alert
			}
			g.World.Credit -= v.Cost
			g.Symptoms[i].Unlocked = true
			return &v, nil
		}
	}
	return nil, errors.New("invalid id")
}

func ApplySymptom(s *model.Symptom, w *model.World) {
	w.Disease.Mortality += s.MortalityBonus
	w.Disease.Transmission += s.TransmissionBonus
}
