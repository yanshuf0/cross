package handle

import (
	"net/http"

	"github.com/gorilla/mux"
)

// InitMux intializes the router with its routes and returns
// a pointer to the mux.
func InitMux(env *Env) *mux.Router {
	// Declare the mux instance.
	m := mux.NewRouter()
	// Declare the api subrouter.
	api := m.PathPrefix("/api").Subrouter()
	// Register coffee machine routes.
	api.HandleFunc("/product/machine", env.coffeeMachines).Methods("GET", "OPTIONS")
	api.HandleFunc("/cross/machine", env.crossSellCoffeeMachines).Methods("GET", "OPTIONS")
	// Register pod routes.
	api.HandleFunc("/product/pod", env.pods).Methods("GET", "OPTIONS")
	api.HandleFunc("/cross/pod", env.crossSellPods).Methods("GET", "OPTIONS")
	// Handle spa routing.
	spa := http.StripPrefix("/", http.FileServer(http.Dir(*env.AssetsDir+"/cross-spa")))
	m.PathPrefix("/").Handler(spa)

	return m
}
