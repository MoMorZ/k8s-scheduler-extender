package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/MoMorZ/k8s-scheduler-extender/controller"

	"github.com/julienschmidt/httprouter"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	router := httprouter.New()
	router.GET("/", controller.Index)
	router.POST("/filter", controller.Predicate)
	router.POST("/prioritize", controller.Prioritize)

	log.Println("Hello I am k8s-scheduler-extender!")
	log.Fatal(http.ListenAndServe(":8888", router))
}
