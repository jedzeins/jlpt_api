package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/jedzeins/jlpt_api/userService/src/services"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	status := services.HealthcheckService()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}
