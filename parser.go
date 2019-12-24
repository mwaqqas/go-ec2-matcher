package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func extractEC2Offer(file string) (o offer, err error) {
	// Read file
	f, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	// parse the offer index
	var oIndex offerIndex
	err = json.Unmarshal([]byte(f), &oIndex)
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, v := range oIndex.Offers {
		if k == "AmazonEC2" {
			return v, err
		}
	}
	return
}
