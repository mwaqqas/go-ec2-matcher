package ec2

import (
	"log"
	"strconv"
)

type EC2 struct {
	WorkloadName string
	CPUCount     string
	RAM          float64
	InstanceType []string
}

func GetEC2Match(csvIn string, region string) (result []EC2, err error) {
	records, err := readCSVFile(csvIn)
	if err != nil {
		log.Fatal(err)
	}

	index, err := extractEC2Product("data/ec2_offer_" + region + ".json")
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		var ramValue float64
		ramValue, err = strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Fatal(err)
		}
		item := EC2{
			WorkloadName: record[0],
			CPUCount:     record[1],
			RAM:          ramValue,
		}

		item.InstanceType = findEC2Instance(item, index.Products)

		result = append(result, item)
	}
	return

	// err = writeCSV(csvOut, result)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}

func findEC2Instance(item EC2, productIndex map[string]ec2) []string {
	var l, u []string
	r := make(map[string]bool)
	for _, product := range productIndex {
		if product.Attributes.Vcpu == item.CPUCount &&
			product.Attributes.Memory.value == item.RAM {
			instance := product.Attributes.InstanceType
			l = append(l, instance)
		}
	}
	for _, val := range l {
		if _, ok := r[val]; !ok {
			r[val] = true
			u = append(u, val)
		}
	}
	return u
}
