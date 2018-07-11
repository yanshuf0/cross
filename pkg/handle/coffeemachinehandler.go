package handle

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// coffeeMachines will return coffee machines based on sizeID
func (env *Env) coffeeMachines(w http.ResponseWriter, r *http.Request) {
	var id int
	var err error
	q := r.URL.Query()
	if q["size_id"] != nil {
		id, err = strconv.Atoi(q["size_id"][0])
		if err != nil {
			http.Error(w, "id is in an invalid format please provide an integer",
				http.StatusBadRequest)
			return
		}
	}
	machines, err := env.DB.Machines(id)
	if err != nil {
		log.Printf("err %v", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(machines)
	if err != nil {
		log.Printf("err %v", err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// crossSellCoffeeMachines will return pods, smallest per flavor that fit
// based on a coffeeMachineID.
func (env *Env) crossSellCoffeeMachines(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	if q["coffee_machine_id"] == nil || len(q["coffee_machine_id"]) > 1 {
		http.Error(w, "please provide one machine id", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(q["coffee_machine_id"][0])
	if err != nil {
		http.Error(w, "id is in an invalid format please provide an integer",
			http.StatusBadRequest)
		return
	}
	pods, err := env.DB.CrossSellMachines(id)
	if err != nil {
		log.Printf("err %v", err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(pods)
	if err != nil {
		log.Printf("err %v", err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
