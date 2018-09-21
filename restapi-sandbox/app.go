package main

import (
	"encoding/json"
	"log"
	"net/http"

	. "github.com/dwdii/go-dwdii/restapi-sandbox/dao"
	. "github.com/dwdii/go-dwdii/restapi-sandbox/models"
	"github.com/gorilla/mux"
)

// Thanks to mlabouardy for examples from https://github.com/mlabouardy/movies-restapi/blob/master/app.go

var dao = PointsDAO{}

func AllPointsEndPoint(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "not implemented yet!")

	res, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		respondWithJson(w, http.StatusOK, res)
	}
}

func CreatePointEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var point Point

	if err := json.NewDecoder(r.Body).Decode(&point); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload"+err.Error())
		return
	}

	params := mux.Vars(r)
	point.UserId = params["userid"]

	err := dao.Insert(point)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/points", AllPointsEndPoint).Methods("GET")
	r.HandleFunc("/points/{userid}", CreatePointEndPoint).Methods("POST")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
