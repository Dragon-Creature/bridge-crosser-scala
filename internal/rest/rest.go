package rest

import (
	"encoding/json"
	"fmt"
	"git.ssns.se/git/frozendragon/bridge-crosser-scala/internal/model"
	"git.ssns.se/git/frozendragon/bridge-crosser-scala/internal/service"
	"github.com/asaskevich/govalidator"
	"io"
	"net/http"
)

type Rest struct {
	Service service.Interfaces
}

func (re *Rest) PostCalculateCrossing(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%+v\n", err)
		RespondError(w, model.Error{
			HttpCode: http.StatusInternalServerError,
			Message:  "unable to read body",
		})
		return
	}
	var request model.CrossingRequest
	err = json.Unmarshal(data, &request)
	if err != nil {
		fmt.Printf("%+v\n", err)
		RespondError(w, model.Error{
			HttpCode: http.StatusBadRequest,
			Message:  "unable to parse body",
		})
		return
	}
	_, err = govalidator.ValidateStruct(request)
	if err != nil {
		RespondError(w, model.Error{
			HttpCode: http.StatusUnprocessableEntity,
			Message:  err.Error(),
		})
		return
	}
	atLeastOneHiker := false
	for _, bridge := range request.Bridges {
		if len(bridge.Hikers) > 0 {
			atLeastOneHiker = true
		}
	}
	if atLeastOneHiker == false {
		RespondError(w, model.Error{
			HttpCode: http.StatusUnprocessableEntity,
			Message:  "you need at least one hiker",
		})
		return
	}

	response := re.Service.CalculateCrossing(request)
	data, err = json.Marshal(response)
	if err != nil {
		fmt.Printf("%+v\n", err)
		RespondError(w, model.Error{
			HttpCode: http.StatusInternalServerError,
			Message:  "unable to create response",
		})
		return
	}
	_, err = w.Write(data)
	if err != nil {
		fmt.Printf("%+v\n", err)
		RespondError(w, model.Error{
			HttpCode: http.StatusInternalServerError,
			Message:  "unable to write response",
		})
		return
	}
}

func RespondError(w http.ResponseWriter, err model.Error) {
	fmt.Printf("%+v\n", err)
	data, callErr := json.Marshal(err)
	if callErr != nil {
		fmt.Printf("%+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(err.HttpCode)
	_, callErr = w.Write(data)
	if callErr != nil {
		fmt.Printf("%+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
