package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	log.SetOutput(os.Stderr)
	//let db spin up before connecting
	time.Sleep(5 * time.Second)
	db := connect()
	err := createSchema(db)
	if err != nil {
		panic(err)
	}
	router := httprouter.New()
	router.GET("/", showLandingPage)
	router.GET("/truetosize", showTrueToSizeAverage)
	router.POST("/truetosize", saveSizeEntry)
	log.Fatal(http.ListenAndServe(":8080", router))
}
