package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mwaqqas/go-ec2-matcher/pkg/ec2"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode("Endpoint is functional")
	if err != nil {
		log.Fatal(err)
	}
}

func EC2MatchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var req ec2.SimpleSearchReq
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Fatal(err)
	}
	result := ec2.SimpleSearch(req)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Fatal(err)
	}

}
