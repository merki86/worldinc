package logic

import "worldinc/app/internal/model"

func BuySymptom(id int, s *[]model.Symptom) *model.Symptom {
	for i, v := range *s {
		if v.ID == id {
			(*s)[i].Unlocked = true
			return &v
		}
	}
	return nil
}

func ApplySymptom(s *model.Symptom, w *model.World) {
	w.Disease.Mortality += s.MortalityBonus
	w.Disease.Transmission += s.TransmissionBonus
}
