package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	. "github.com/jafetbntz/candy_shop/dao"
	"gopkg.in/mgo.v2/bson"

	. "github.com/jafetbntz/candy_shop/models"

	"github.com/gorilla/mux"
)

var dao = CandiesDAO{}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "service running...")
}

func AllCandies(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Aún no implementado")
}

func FindCandy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Aún no implementado")
}

func CreateCandy(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var candy Candy
	if err := json.NewDecoder(r.Body).Decode(&candy); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	candy.Id = bson.NewObjectId()

	if err := dao.Insert(candy); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, candy)
}

func UpdateCandy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Aún no implementado")
}

func DeleteCandy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Aún no implementado")
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	fmt.Println(msg)
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {

	dao.Server = "localhost"
	dao.Database = "can_db"
	dao.Connect()
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", Index).Methods("GET")
	r.HandleFunc("/candy", AllCandies).Methods("GET")
	r.HandleFunc("/candy", CreateCandy).Methods("POST")
	r.HandleFunc("/candy", UpdateCandy).Methods("PUT")
	r.HandleFunc("/candy", DeleteCandy).Methods("DELETE")
	r.HandleFunc("/candy/{id}", FindCandy).Methods("GET")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}

}
