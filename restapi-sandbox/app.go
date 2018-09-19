package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	. "github.com/dwdii/go-dwdii/restapi-sandbox/dao"
	"github.com/gorilla/mux"
)

var dao = PointsDAO{}

func AllPointsEndPoint(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "not implemented yet!")

	res, err := dao.FindAll()
	if err != nil {
		fmt.Fprintln(w, err)
	} else {
		respondWithJson(w, http.StatusOK, res)
	}
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
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
