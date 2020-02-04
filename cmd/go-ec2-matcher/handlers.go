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
	var request []ec2.SimpleSearchReq
	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Fatal(err)
	}

	// list of maps whose keys are string
	// the values are a list of maps whose keys are string
	// and values are ec2.ResultInstance
	var rs []map[string][]map[string]ec2.ResultInstance
	for _, item := range request {
		m := make(map[string][]map[string]ec2.ResultInstance)
		result := ec2.SimpleSearch(item)
		m[item.WorkloadName] = result
		rs = append(rs, m)
	}

	err = json.NewEncoder(w).Encode(rs)
	if err != nil {
		log.Fatal(err)
	}

}

func EC2PriceHandler(w http.ResponseWriter, r *http.Request) {

}
