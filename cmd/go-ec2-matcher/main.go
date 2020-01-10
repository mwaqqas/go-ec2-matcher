package main

import (
	"log"

	"github.com/mwaqqas/go-ec2-matcher/pkg/ec2"
)

func main() {
	err := ec2.GetEC2Prices(
		"sample_data/ec2-requirements.csv",
		"sample_data/ec2-prices-result.csv",
		"me-south-1",
	)
	if err != nil {
		log.Fatal(err)
	}

}
