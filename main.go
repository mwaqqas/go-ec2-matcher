package main

import (
	"log"
)

func main() {
	err := GetEC2Prices(
		"sample_data/ec2-requirements.csv",
		"sample_data/ec2-prices-result.csv",
		"me-south-1",
	)
	if err != nil {
		log.Fatal(err)
	}

}
