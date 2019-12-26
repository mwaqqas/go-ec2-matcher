package main

import (
	"fmt"
	// "os"
)

// ec2instance.info = https://raw.githubusercontent.com/powdahound/ec2instances.info/master/www/instances.json

func main() {
	// Create Data Directory
	// if _, err := os.Stat(dataDirPath); err != nil {
	// 	os.MkdirAll(dataDirPath, 0755)
	// }

	// downloadFile(offerIndexURL, offerIndexPath, false)

	// EC2Offer, err := extractEC2Offer(offerIndexPath)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// EC2RegionIndexURL := baseURL + EC2Offer.CurrentRegionIndexURL
	// downloadFile(EC2RegionIndexURL, regionIndexPath, false)
	// downloadEC2RegionOffers(regionIndexPath)
	// index, err := extractEC2Product("sample_data/me-south-1_ec2.json")
	index, err := extractEC2Product("sample_data/ec2_offer_me-south-1.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range index.Products {
		fmt.Printf(
			"{InstanceType: %s, vCPU: %s, RAM: %.2f %s, OS: %s, InstalledSW: %s, Bandwidth: %s}\n",
			v.Attributes.InstanceType,
			v.Attributes.Vcpu,
			v.Attributes.Memory.value,
			v.Attributes.Memory.unit,
			v.Attributes.OperatingSystem,
			v.Attributes.PreInstalledSw,
			v.Attributes.NetworkPerformance,
		)
	}
}
