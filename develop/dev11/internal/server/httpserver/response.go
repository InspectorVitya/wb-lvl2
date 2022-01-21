package httpserver

import (
	"encoding/json"
	"github.com/inspectorvitya/wb-lvl2/develop/dev11/internal/model"
	"net/http"
)

type errorResponse struct {
	Message string `json:"error"`
}

func writeJSONResponse(w http.ResponseWriter, statusCode int, msg string, isErr bool) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if isErr {
		msgWithErr := errorResponse{Message: msg}
		json.NewEncoder(w).Encode(msgWithErr)
		return
	}
	data := struct {
		Message string `json:"Result"`
	}{Message: msg}
	json.NewEncoder(w).Encode(data)
}

func writeJSONEvents(w http.ResponseWriter, statusCode int, events []model.Event) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	data := struct {
		Events []model.Event `json:"Result"`
	}{Events: events}
	json.NewEncoder(w).Encode(data)
}
