package logic

import "worldinc/app/internal/model"

func Buy(id int, s *[]model.Symptom) {
	for i, v := range *s {
		if v.ID == id {
			(*s)[i].Unlocked = true
		}
	}
}
