package util

import (
	"database/sql"
	"encoding/json"
	"go-multitenancy/model"
	"net/http"
)

func SendServerError(w http.ResponseWriter, err error) {
	var e model.Error
	e.Message = err.Error()
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(e)
}

func SendBadRequest(w http.ResponseWriter, err error) {
	var e model.Error
	e.Message = err.Error()
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(e)

}

func SendNotFound(w http.ResponseWriter, err error) {
	var e model.Error
	e.Message = err.Error()
	w.WriteHeader(http.StatusNotFound)
	_ = json.NewEncoder(w).Encode(e)
}

func SendSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}

func SendResult(w http.ResponseWriter, data interface{}, err error) {
	if err != nil {
		if err == sql.ErrNoRows {
			SendNotFound(w, err)
		} else {
			SendServerError(w, err)
		}
		return
	}

	SendSuccess(w, data)
}

func ResultData(data interface{}, z interface{}, err error) (interface{}, error) {
	if err != nil {
		return z, err
	}

	return data, nil
}
