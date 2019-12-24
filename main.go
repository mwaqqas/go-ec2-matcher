package main

import (
	"fmt"
	"os"
)

// ec2instance.info = https://raw.githubusercontent.com/powdahound/ec2instances.info/master/www/instances.json

func main() {
	// Create Data Directory
	if _, err := os.Stat(dataDirPath); err != nil {
		os.MkdirAll(dataDirPath, 0755)
	}

	downloadFile(offerIndexURL, offerIndexPath, false)

	EC2Offer, err := extractEC2Offer(offerIndexPath)
	if err != nil {
		fmt.Println(err)
	}
	EC2RegionIndexURL := baseURL + EC2Offer.CurrentRegionIndexURL
	downloadFile(EC2RegionIndexURL, regionIndexPath, false)
	downloadEC2RegionOffers(regionIndexPath)
}
