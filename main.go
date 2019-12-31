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
	index, err := extractEC2Product("sample_data/bh-c5.xlarge.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, prodV := range index.Products {
		var date string
		var pd map[string]priceDimensions
		for k, v := range index.Terms.OnDemand {
			if k == prodV.Sku {
				refKey := prodV.Sku + onDemandOfferCode
				date = v[refKey].EffectiveDate
				pd = v[refKey].PriceDimensions
				for _, v := range pd {
					fmt.Printf(
						"{SKU: %s, InstanceType: %s, vCPU: %s, RAM: %.2f %s,  Bandwidth: %s, OS: %s, InstalledSW: %s, EffectiveDate: %s}\n",
						prodV.Sku,
						prodV.Attributes.InstanceType,
						prodV.Attributes.Vcpu,
						prodV.Attributes.Memory.value,
						prodV.Attributes.Memory.unit,
						prodV.Attributes.NetworkPerformance,
						prodV.Attributes.OperatingSystem,
						prodV.Attributes.PreInstalledSw,
						date,
					)
					fmt.Printf(
						"Description: %s, Price: %s, LeaseContractLength: %s, OfferingClass: %s, PurchaseOption: %s\n",
						v.Description,
						v.PricePerUnit.USD,
						v.TermAttributes.LeaseContractLength,
						v.TermAttributes.OfferingClass,
						v.TermAttributes.PurchaseOption,
					)
				}
			}
		}
		fmt.Println("-------------------------")
	}

}
