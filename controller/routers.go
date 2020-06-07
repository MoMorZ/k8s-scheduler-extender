package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Well~Welcome to k8s-scheduler-extender!")
}

func Predicate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	checkBody(w, r)
	var buf bytes.Buffer
	var extenderArgs schedulerapi.ExtenderArgs
	var extenderFilterResult *schedulerapi.ExtenderFilterResult
	body := io.TeeReader(r.Body, &buf)
	if err := json.NewDecoder(body).Decode(&extenderArgs); err != nil {
		extenderFilterResult = &schedulerapi.ExtenderFilterResult{
			Nodes:       nil,
			FailedNodes: nil,
			Error:       err.Error(),
		}
	} else {
		extenderFilterResult = filter(extenderArgs)
	}
	if response, err := json.Marshal(extenderFilterResult); err != nil {
		log.Fatalln(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write(response); err != nil {
			log.Println("[Error]Write fail...TaT")
		}
	}
}

func Prioritize(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	checkBody(w, r)
	var buf bytes.Buffer
	body := io.TeeReader(r.Body, &buf)
	var extenderArgs schedulerapi.ExtenderArgs
	var hostPriorityList *schedulerapi.HostPriorityList
	if err := json.NewDecoder(body).Decode(&extenderArgs); err != nil {
		hostPriorityList = &schedulerapi.HostPriorityList{}
	} else {
		hostPriorityList = prioritize(extenderArgs)
	}

	if response, err := json.Marshal(hostPriorityList); err != nil {
		log.Fatalln(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write(response); err != nil {
			log.Println("[Error]Write fail...TaT")
		}
	}
}

// func Bind(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {}

func checkBody(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "[Error]So bad, your body is empty...QAQ", 400)
	}
}
