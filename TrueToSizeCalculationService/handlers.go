package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	validator "gopkg.in/go-playground/validator.v9"
)

func showLandingPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello, web")
}

//returns JSON of the
func showTrueToSizeAverage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//fmt.Fprint(w, "Welcome!\n")
	db := connect()
	defer db.Close()

	var trueToSizes []TrueToSize
	err := db.Model(&trueToSizes).Select()
	if err != nil {
		panic(err)
	}

	total := 0
	for _, ttsRow := range trueToSizes {
		total += ttsRow.Size
	}
	numberOfEntries := len(trueToSizes)
	var result = TrueToSizeResult{}
	result.AverageTrueToSize = float32(total) / float32(numberOfEntries)
	responseJSON(w, result)
}

//saves a size received on the POST to the url
func saveSizeEntry(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New()

	// Read posted data
	truetoSize := &TrueToSize{}
	err = json.Unmarshal(data, truetoSize)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = validate.Struct(truetoSize)
	if err != nil {
		validationMessage := ""
		errs := err.(validator.ValidationErrors)
		validationMessage = fmt.Sprintf("%s", errs)
		responseError(w, validationMessage, http.StatusBadRequest)
		return
	}

	// Insert new post
	db := connect()
	defer db.Close()
	err = db.Insert(truetoSize)
	if err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseJSON(w, truetoSize)
}
