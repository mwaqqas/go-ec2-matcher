package main

import (
	"fmt"
	"os"
)

// ec2instance.info = https://raw.githubusercontent.com/powdahound/ec2instances.info/master/www/instances.json

func oldmain() {
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
	err = downloadEC2RegionOffers(regionIndexPath)
	if err != nil {
		fmt.Println("Download file error.")
	}
	index, err := extractEC2Product("data/ec2_offer_me-south-1.json")
	// index, err := extractEC2Product("sample_data/bh-c5.xlarge.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// ON DEMAND
	// for _, prodV := range index.Products {
	// 	var date string
	// 	var pd map[string]priceDimensions
	// 	for k, v := range index.Terms.OnDemand {
	// 		if k == prodV.Sku {
	// 			refKey := prodV.Sku + onDemandOfferCode
	// 			date = v[refKey].EffectiveDate
	// 			pd = v[refKey].PriceDimensions
	// 			for _, v := range pd {
	// 				fmt.Printf(
	// 					"{SKU: %s, InstanceType: %s, vCPU: %s, RAM: %.2f %s,  Bandwidth: %s, OS: %s, InstalledSW: %s, EffectiveDate: %s}\n",
	// 					prodV.Sku,
	// 					prodV.Attributes.InstanceType,
	// 					prodV.Attributes.Vcpu,
	// 					prodV.Attributes.Memory.value,
	// 					prodV.Attributes.Memory.unit,
	// 					prodV.Attributes.NetworkPerformance,
	// 					prodV.Attributes.OperatingSystem,
	// 					prodV.Attributes.PreInstalledSw,
	// 					date,
	// 				)
	// 				fmt.Printf(
	// 					"Description: %s, Price: %s, LeaseContractLength: %s, OfferingClass: %s, PurchaseOption: %s\n",
	// 					v.Description,
	// 					v.PricePerUnit.USD,
	// 					v.TermAttributes.LeaseContractLength,
	// 					v.TermAttributes.OfferingClass,
	// 					v.TermAttributes.PurchaseOption,
	// 				)
	// 			}
	// 		}
	// 	}
	// 	fmt.Println("-------------------------")
	// }

	// RESERVED
	// for sku, offerTerms := range index.Terms.Reserved {
	// 	for offerSKU, offer := range offerTerms {
	// 		fmt.Printf("ProdSKU: %s, offerSKU: %s, EffectiveDate: %s\n", sku, offerSKU, offer.EffectiveDate)
	// 	}
	// }

	EC2SKU := "WY32JPCE2DYBBWNK"
	LeaseContractLength := "3yr"
	PurchaseOption := "No Upfront"
	OfferingClass := "convertible"

	fmt.Println("\nPRODUCT DETAILS")
	for productSKU, product := range index.Products {
		if productSKU == EC2SKU {
			fmt.Printf(
				"productSKU: %s, InstanceType: %s, vCPU: %s, Memory: %.2f %s, OS: %s, PreInstalledSW: %s\n",
				productSKU,
				product.Attributes.InstanceType,
				product.Attributes.Vcpu,
				product.Attributes.Memory.value,
				product.Attributes.Memory.unit,
				product.Attributes.OperatingSystem,
				product.Attributes.PreInstalledSw,
			)
		}
	}

	fmt.Println("\nONDEMAND")
	for productSKU, onDemandOfferTermMap := range index.Terms.OnDemand {
		if productSKU == EC2SKU {
			for offerTermSKU, offerTerm := range onDemandOfferTermMap {
				for pdSKU, pd := range offerTerm.PriceDimensions {
					fmt.Printf(
						"productSKU: %s, offerTermSKU: %s, pdSKU: %s,\n"+
							"Description: %s, pricePerUnit: %s\n",
						productSKU,
						offerTermSKU,
						pdSKU,
						pd.Description,
						pd.PricePerUnit.USD,
					)
				}
			}
		}
	}

	fmt.Println("\nRESERVED")
	for productSKU, onDemandOfferTermMap := range index.Terms.Reserved {
		if productSKU == EC2SKU {
			for offerTermSKU, offerTerm := range onDemandOfferTermMap {
				if offerTerm.TermAttributes.LeaseContractLength == LeaseContractLength &&
					offerTerm.TermAttributes.OfferingClass == OfferingClass &&
					offerTerm.TermAttributes.PurchaseOption == PurchaseOption {
					for pdSKU, pd := range offerTerm.PriceDimensions {
						fmt.Printf(
							"productSKU: %s, offerTermSKU: %s, pdSKU: %s, description: %s, pricePerUnit: %s\n"+
								"TermAttributes: {LeaseContractLength: %s, PurchaseOption: %s, OfferingClass: %s}\n"+
								"---------\n",
							productSKU,
							offerTermSKU,
							pdSKU,
							pd.Description,
							pd.PricePerUnit.USD,
							offerTerm.TermAttributes.LeaseContractLength,
							offerTerm.TermAttributes.PurchaseOption,
							offerTerm.TermAttributes.OfferingClass,
						)
					}
				}
			}
		}
	}

}
