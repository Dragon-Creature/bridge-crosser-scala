package rest

import (
	"encoding/json"
	"fmt"
	"git.ssns.se/git/frozendragon/bridge-crosser-scala/internal/model"
	"git.ssns.se/git/frozendragon/bridge-crosser-scala/internal/service"
	"io"
	"net/http"
)

func PostCalculateCrossing(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var request model.CrossingRequest
	err = json.Unmarshal(data, &request)
	if err != nil {
		fmt.Printf("%+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := service.CalculateCrossing(request)
	if err != nil {
		fmt.Printf("%+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err = json.Marshal(response)
	if err != nil {
		fmt.Printf("%+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
}
