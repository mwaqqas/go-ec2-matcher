package ec2

import (
	"encoding/csv"
	"log"
	"os"
)

// import (
// 	"encoding/csv"
// 	"log"
// 	"os"
//
//
//
// )

type EC2Item struct {
	WorkloadName        string
	InstanceType        string
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

// func GetEC2Prices(csvIn string, csvOut string, region string) (err error) {
// 	records, err := readCSVFile(csvIn)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var result []EC2Item
// 	header := makeHeaderRow()
// 	result = append(result, header)

// 	index, err := extractEC2Product("data/ec2_offer_" + region + ".json")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, record := range records {
// 		item := makeItemFromCSVRow(record)
// 		prodSKU := findEC2ItemSKU(item, index.Products)
// 		switch item.PaymentModel {
// 		case "OnDemand":
// 			r, found := findOnDemandOffer(prodSKU, item, index.Terms.OnDemand)
// 			if found {
// 				result = append(result, r)
// 			}
// 		case "Reserved":
// 			r, found := findReservedOffer(prodSKU, item, index.Terms.Reserved)
// 			if found {
// 				result = append(result, r)
// 			}
// 		}
// 	}

// 	err = writeCSV(csvOut, result)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return

// }

func readCSVFile(filename string) (rows [][]string, err error) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)

	// skip the header row
	if _, err := r.Read(); err != nil {
		log.Fatal(err)
	}

	// process the rest of the records
	rows, err = r.ReadAll()
	return
}

func makeHeaderRow() EC2Item {
	return EC2Item{
		WorkloadName:        "WorkloadName",
		InstanceType:        "InstanceType",
		Environment:         "Environment",
		OS:                  "OS",
		PreInstalledSw:      "PreInstalledSw",
		Tenancy:             "Tenancy",
		PaymentModel:        "PaymentModel",
		LeaseContractLength: "LeaseContractLength",
		OfferingClass:       "OfferingClass",
		PurchaseOption:      "PurchaseOption",
		UpfrontFee:          "UpfrontFee (USD)",
		HourlyFee:           "HourlyFee (USD)",
	}
}

// func makeItemFromCSVRow(record []string) EC2Item {
// 	return EC2Item{
// 		WorkloadName:        record[0],
// 		InstanceType:        record[1],
// 		Environment:         record[2],
// 		OS:                  record[3],
// 		PreInstalledSw:      record[4],
// 		Tenancy:             record[5],
// 		PaymentModel:        record[6],
// 		LeaseContractLength: record[7],
// 		OfferingClass:       record[8],
// 		PurchaseOption:      record[9],
// 	}
// }

// func findEC2ItemSKU(item EC2Item, productIndex map[string]ec2) (sku string) {
// 	for _, product := range productIndex {
// 		if product.Attributes.InstanceType == item.InstanceType &&
// 			product.Attributes.PreInstalledSw == item.PreInstalledSw &&
// 			product.Attributes.OperatingSystem == item.OS &&
// 			product.Attributes.Tenancy == item.Tenancy &&
// 			product.Attributes.Capacitystatus == "Used" {
// 			return product.Sku
// 		}
// 	}
// 	return //empty
// }

// func findOnDemandOffer(sku string, item EC2Item, terms map[string]map[string]offerTerm) (out EC2Item, found bool) {
// 	for k, v := range terms {
// 		if k == sku {
// 			for _, term := range v {
// 				for _, pd := range term.PriceDimensions {
// 					item.HourlyFee = pd.PricePerUnit.USD
// 					return item, true
// 				}
// 			}
// 		}
// 	}
// 	return item, false
// }

// func findReservedOffer(sku string, item EC2Item, terms map[string]map[string]offerTerm) (out EC2Item, found bool) {
// 	for k, v := range terms {
// 		if k == sku {
// 			for _, term := range v {
// 				if term.TermAttributes.LeaseContractLength == item.LeaseContractLength &&
// 					term.TermAttributes.OfferingClass == item.OfferingClass &&
// 					term.TermAttributes.PurchaseOption == item.PurchaseOption {
// 					for _, pd := range term.PriceDimensions {
// 						if pd.Description == "Upfront Fee" {
// 							item.UpfrontFee = pd.PricePerUnit.USD
// 						} else {
// 							item.HourlyFee = pd.PricePerUnit.USD
// 						}
// 					}
// 					return item, true
// 				}
// 			}
// 		}
// 	}
// 	return item, false
// }

// func writeCSV(filename string, contents []EC2Item) (err error) {
// 	f, err := os.Create(filename)
// 	if err != nil {
// 		log.Fatal(err)
// 		return err
// 	}
// 	defer f.Close()

// 	writer := csv.NewWriter(f)
// 	defer writer.Flush()

// 	for _, item := range contents {
// 		s := []string{
// 			item.WorkloadName,
// 			item.InstanceType,
// 			item.Environment,
// 			item.OS,
// 			item.PreInstalledSw,
// 			item.Tenancy,
// 			item.PaymentModel,
// 			item.LeaseContractLength,
// 			item.OfferingClass,
// 			item.PurchaseOption,
// 			item.UpfrontFee,
// 			item.HourlyFee,
// 		}

// 		err = writer.Write(s)
// 		if err != nil {
// 			log.Fatal(err)
// 			return
// 		}
// 	}
// 	return
// }
