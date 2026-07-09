package api

import (
	"encoding/json"
	"net/http"

	"github.com/zinrai/debiface-gen/config"
)

type validatable interface {
	Validate() error
}

// handleConfig decodes a config of type T from the request body, validates it,
// and returns the generated interfaces(5) stanza as JSON ({"config": ...}) so
// the client (e.g. an IPAM-integrated agent) can parse it and apply it.
func handleConfig[T validatable](w http.ResponseWriter, r *http.Request, generate func(T) string) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var cfg T
	if err := json.NewDecoder(r.Body).Decode(&cfg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := cfg.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"config": generate(cfg)})
}

func HandleBonding(w http.ResponseWriter, r *http.Request) {
	handleConfig(w, r, config.GenerateBondingConfig)
}

func HandleDSR(w http.ResponseWriter, r *http.Request) {
	handleConfig(w, r, config.GenerateDSRConfig)
}

func HandleStandard(w http.ResponseWriter, r *http.Request) {
	handleConfig(w, r, config.GenerateStandardConfig)
}

func HandleBridge(w http.ResponseWriter, r *http.Request) {
	handleConfig(w, r, config.GenerateBridgeConfig)
}
