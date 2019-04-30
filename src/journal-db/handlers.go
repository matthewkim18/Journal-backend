package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func CreateJournal(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	journal := Journal{}
	err := json.NewDecoder(r.Body).Decode(&journal)
	if err != nil {
		writeErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		return
	}

	journal.ID = bson.NewObjectId()

	err1 := db.C("journals").Insert(&journal)
	if err1 != nil {
		fmt.Printf(err1.Error())
	}
	writeOKResponse(w, journal)
}
func DeleteJournal(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	journal := Journal{}
	id := params.ByName("id")
	err := db.C("journals").Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		writeErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		return
	}
	writeOKResponse(w, journal)
}
func UpdateJournal(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	journal := Journal{}
	err := json.NewDecoder(r.Body).Decode(&journal)
	if err != nil {
		writeErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		return
	}
	id := params.ByName("id")
	err1 := db.C("journals").Update(bson.M{"_id": bson.ObjectIdHex(id)}, &journal)
	//journal_list[journal.ID] = journal
	if err1 != nil {
		fmt.Printf(err1.Error())
	}
	writeOKResponse(w, journal)
}

func FindJournal(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	journal := []Journal{}

	err := db.C("journals").Find(bson.M{}).All(&journal)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("%v", journal)
	writeOKResponse(w, journal)
}

func ShowJournal(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	journal := Journal{}
	err := db.C("journals").Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&journal)
	if err != nil {
		writeErrorResponse(w, http.StatusNotFound, "Record Not Found")
		return
	}

	writeOKResponse(w, journal)

}

// Writes the response as a standard JSON response with StatusOK
func writeOKResponse(w http.ResponseWriter, m interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&m); err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
	}
}

// Writes the error response as a Standard API JSON response with a response code
func writeErrorResponse(w http.ResponseWriter, errorCode int, errorMsg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errorCode)
	json.
		NewEncoder(w).
		Encode(&JsonErrorResponse{Error: &ApiError{Status: errorCode, Title: errorMsg}})
}

//Populates a model from the params in the Handler
func populateModelFromHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params, model interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return err
	}
	if err := r.Body.Close(); err != nil {
		return err
	}
	if err := json.Unmarshal(body, model); err != nil {
		return err
	}
	return nil
}
