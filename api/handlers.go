package api

import (
	"encoding/json"
	"net/http"

	"github.com/zinrai/debiface-gen/config"
)

func HandleBonding(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var cfg config.BondingConfig
	err := json.NewDecoder(r.Body).Decode(&cfg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := config.GenerateBondingConfig(cfg)
	json.NewEncoder(w).Encode(map[string]string{"config": result})
}

func HandleDSR(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var cfg config.DSRConfig
	err := json.NewDecoder(r.Body).Decode(&cfg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := config.GenerateDSRConfig(cfg)
	json.NewEncoder(w).Encode(map[string]string{"config": result})
}

func HandleStandard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var cfg config.StandardConfig
	err := json.NewDecoder(r.Body).Decode(&cfg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := config.GenerateStandardConfig(cfg)
	json.NewEncoder(w).Encode(map[string]string{"config": result})
}
