package main

import (
	"encoding/csv"
	"log"
	"os"
)

type EC2Req struct {
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

func GetEC2Prices(csvIn string, csvOut string, region string) (err error) {
	f, err := os.Open(csvIn)
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
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	var out [][]string
	out = append(out, []string{
		"WorkloadName",
		"InstanceType",
		"Environment",
		"OS",
		"PreInstalledSw",
		"Tenancy",
		"PaymentModel",
		"LeaseContractLength",
		"OfferingClass",
		"PurchaseOption",
		"ProductSKU",
		"Upfront Fee",
		"Hourly Price",
	})

	index, err := extractEC2Product("data/ec2_offer_" + region + ".json")
	if err != nil {
		log.Fatal(err)
	}
	for _, record := range records {
		req := EC2Req{
			WorkloadName:        record[0],
			InstanceType:        record[1],
			Environment:         record[2],
			OS:                  record[3],
			PreInstalledSw:      record[4],
			Tenancy:             record[5],
			PaymentModel:        record[6],
			LeaseContractLength: record[7],
			OfferingClass:       record[8],
			PurchaseOption:      record[9],
		}
		for prodK, prodV := range index.Products {
			if prodV.Attributes.InstanceType == req.InstanceType &&
				prodV.Attributes.PreInstalledSw == req.PreInstalledSw &&
				prodV.Attributes.OperatingSystem == req.OS &&
				prodV.Attributes.Tenancy == req.Tenancy &&
				prodV.Attributes.Capacitystatus == "Used" {
				switch req.PaymentModel {
				case "OnDemand":
					for prodSKU, termMap := range index.Terms.OnDemand {
						if prodSKU == prodK {
							for _, offerTerm := range termMap {
								for _, pd := range offerTerm.PriceDimensions {
									out = append(out, []string{
										req.WorkloadName,
										req.InstanceType,
										req.Environment,
										req.OS,
										req.PreInstalledSw,
										req.Tenancy,
										req.PaymentModel,
										req.LeaseContractLength,
										req.OfferingClass,
										req.PurchaseOption,
										prodK,
										"", //upfront fee
										pd.PricePerUnit.USD,
									})
								}
							}
						}
					}
				case "Reserved":
					for prodSKU, termMap := range index.Terms.Reserved {
						if prodSKU == prodK {
							for _, offerTerm := range termMap {
								if offerTerm.TermAttributes.LeaseContractLength == req.LeaseContractLength &&
									offerTerm.TermAttributes.OfferingClass == req.OfferingClass &&
									offerTerm.TermAttributes.PurchaseOption == req.PurchaseOption {
									for _, pd := range offerTerm.PriceDimensions {
										if pd.Description == "Upfront Fee" {
											out = append(out, []string{
												req.WorkloadName,
												req.InstanceType,
												req.Environment,
												req.OS,
												req.PreInstalledSw,
												req.Tenancy,
												req.PaymentModel,
												req.LeaseContractLength,
												req.OfferingClass,
												req.PurchaseOption,
												prodK,
												pd.PricePerUnit.USD, //upfront fee
												"",
											})
										} else {
											out = append(out, []string{
												req.WorkloadName,
												req.InstanceType,
												req.Environment,
												req.OS,
												req.PreInstalledSw,
												req.Tenancy,
												req.PaymentModel,
												req.LeaseContractLength,
												req.OfferingClass,
												req.PurchaseOption,
												prodK,
												"", //upfront fee
												pd.PricePerUnit.USD,
											})
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	f, err = os.Create(csvOut)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	for _, value := range out {
		err := writer.Write(value)
		if err != nil {
			log.Fatal(err)
		}
	}

	return

}
