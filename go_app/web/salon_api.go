package web

import (
	"coifResa"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateSalon() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		salon := &coifResa.SalonItem{}

		err := json.NewDecoder(r.Body).Decode(salon)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = h.Store.CreateSalon(salon)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(struct {
			Status  string              `json:"status"`
			Message string              `json:"message"`
			Salon   *coifResa.SalonItem `json:"salon"`
		}{
			Status:  "success",
			Message: "Salon créé avec succès",
			Salon:   salon,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}