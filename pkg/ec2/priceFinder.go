package ec2

import (
	"fmt"
	"log"
)

type EC2Item struct {
	WorkloadName        string
	InstanceType        string
	Region              string
	Environment         string
	OS                  string
	PreInstalledSw      string
	Tenancy             string
	PaymentModel        string
	LeaseContractLength string
	OfferingClass       string
	PurchaseOption      string
	UpfrontFee          string
	HourlyFee           string
}

type PriceRequest []map[string][]EC2Item
type PriceResponse []map[string][]EC2Item

func GetEC2Prices(request PriceRequest) (results PriceResponse, err error) {
	for _, regionMap := range request {
		for region, items := range regionMap {
			index, err := extractEC2Product("data/ec2_offer_" + region + ".json")
			if err != nil {
				log.Fatal(err)
			}
			pricePerRegionMap := make(map[string][]EC2Item)

			for _, item := range items {
				prodSKU := findEC2ItemSKU(item, index.Products)
				switch item.PaymentModel {

				case "OnDemand":
					r, found := findOnDemandOffer(prodSKU, item, index.Terms.OnDemand)
					if found {
						pricePerRegionMap[region] = append(pricePerRegionMap[region], r)
					}
				case "Reserved":
					r, found := findReservedOffer(prodSKU, item, index.Terms.Reserved)
					if found {
						pricePerRegionMap[region] = append(pricePerRegionMap[region], r)
					}
				}
			}
			results = append(results, pricePerRegionMap)
		}
	}

	fmt.Println(results)
	return
}

func findEC2ItemSKU(item EC2Item, productIndex map[string]Product) (sku string) {
	for _, product := range productIndex {
		if product.Attributes.InstanceType == item.InstanceType &&
			product.Attributes.PreInstalledSw == item.PreInstalledSw &&
			product.Attributes.OperatingSystem == item.OS &&
			product.Attributes.Tenancy == item.Tenancy &&
			product.Attributes.Capacitystatus == "Used" {
			return product.Sku
		}
	}
	return //empty
}

func findOnDemandOffer(sku string, item EC2Item, terms map[string]map[string]OfferTerm) (out EC2Item, found bool) {
	for k, v := range terms {
		if k == sku {
			for _, term := range v {
				for _, pd := range term.PriceDimensions {
					item.HourlyFee = pd.PricePerUnit.USD
					return item, true
				}
			}
		}
	}
	return item, false
}

func findReservedOffer(sku string, item EC2Item, terms map[string]map[string]OfferTerm) (out EC2Item, found bool) {
	for k, v := range terms {
		if k == sku {
			for _, term := range v {
				if term.TermAttributes.LeaseContractLength == item.LeaseContractLength &&
					term.TermAttributes.OfferingClass == item.OfferingClass &&
					term.TermAttributes.PurchaseOption == item.PurchaseOption {
					for _, pd := range term.PriceDimensions {
						if pd.Description == "Upfront Fee" {
							item.UpfrontFee = pd.PricePerUnit.USD
						} else {
							item.HourlyFee = pd.PricePerUnit.USD
						}
					}
					return item, true
				}
			}
		}
	}
	return item, false
}
