package handle

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/yanshuf0/cross/pkg/data"
)

// pods returns pods which can be filtered by size and/or flavor, accepted
// parameters are flavorID and sizeID.
func (env *Env) pods(w http.ResponseWriter, r *http.Request) {
	var flavorID int
	var sizeID int
	var err error
	var pods []data.Pod

	q := r.URL.Query()
	if q["flavor_id"] != nil {
		flavorID, err = strconv.Atoi(q["flavor_id"][0])
		if err != nil {
			http.Error(w, "flavor_id is in an invalid format please provide an integer",
				http.StatusBadRequest)
			return
		}
	}
	if q["size_id"] != nil {
		sizeID, err = strconv.Atoi(q["size_id"][0])
		if err != nil {
			http.Error(w, "size_id is in an invalid format please provide an integer",
				http.StatusBadRequest)
			return
		}
	}
	pods, err = env.DB.Pods(flavorID, sizeID)
	if err != nil {
		log.Printf("err: %v", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
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

// crossSellPods will return pods based on a podId, it will return the other
// quantities available for that pod size / flavor.
func (env *Env) crossSellPods(w http.ResponseWriter, r *http.Request) {
	var pods []data.Pod
	var podID int
	var err error
	q := r.URL.Query()
	if q["pod_id"] == nil {
		http.Error(w, "please provide a pod_id", http.StatusBadRequest)
		return
	}
	podID, err = strconv.Atoi(q["pod_id"][0])
	if err != nil {
		http.Error(w, "pod_id is in an invalid format please provide an integer",
			http.StatusBadRequest)
		return
	}
	pods, err = env.DB.CrossSellPods(podID)
	if err != nil {
		log.Printf("err %v", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
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
