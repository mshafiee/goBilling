/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 */

package main

import (
	"log"
	"net/http"
	"github.com/mshafiee/goBilling/model"
)

func main() {
	log.Printf("Server started")

	router := model.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
