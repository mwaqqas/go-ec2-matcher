package ec2

import (
	"log"
)

// type EC2 struct {
// 	WorkloadName string
// 	CPUCount     int
// 	RAM          float64
// 	InstanceType []string
// }

// GetEC2Match : comment
// func GetEC2Match(csvIn string, region string, CPUFuzzFactor int, RAMFuzzFactor float64) (result []EC2, err error) {
// 	records, err := readCSVFile(csvIn)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	index, err := extractEC2Product("data/ec2_offer_" + region + ".json")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, record := range records {
// 		var ramValue float64
// 		var cpuCountVal int
// 		ramValue, err = strconv.ParseFloat(record[2], 64)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		cpuCountVal, err = strconv.Atoi(record[1])
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		item := EC2{
// 			WorkloadName: record[0],
// 			CPUCount:     cpuCountVal,
// 			RAM:          ramValue,
// 		}

// 		item.InstanceType = findEC2Instance(item, CPUFuzzFactor, RAMFuzzFactor, index.Products)

// 		result = append(result, item)
// 	}
// 	return
// }

// func findEC2Instance(item EC2, CPUFuzzMatchFactor int, RAMFuzzMatchFactor float64, productIndex map[string]ec2) []string {
// 	var l, u []string
// 	r := make(map[string]bool)
// 	for _, product := range productIndex {
// 		if product.Attributes.Vcpu == item.CPUCount &&
// 			product.Attributes.Memory.value == item.RAM {
// 			instance := product.Attributes.InstanceType
// 			l = append(l, instance)
// 		}
// 	}
// 	if len(l) < 1 {
// 		for _, product := range productIndex {
// 			CPUMax := item.CPUCount + CPUFuzzMatchFactor
// 			CPUMin := item.CPUCount - CPUFuzzMatchFactor
// 			RAMMax := item.RAM + RAMFuzzMatchFactor
// 			RAMMin := item.RAM - RAMFuzzMatchFactor

// 			if (CPUMin <= product.Attributes.Vcpu && product.Attributes.Vcpu <= CPUMax) &&
// 				(RAMMin <= product.Attributes.Memory.value && product.Attributes.Memory.value <= RAMMax) {
// 				instance := product.Attributes.InstanceType
// 				l = append(l, instance)
// 			}
// 		}
// 	}

// 	for _, val := range l {
// 		if _, ok := r[val]; !ok {
// 			r[val] = true
// 			u = append(u, val)
// 		}
// 	}
// 	return u
// }

type reqInstance struct {
	CPU int
	RAM float64
}

// ByCPUAndRAM : comment
func ByCPUAndRAM(
	region string,
	reqCPU int,
	reqRAM float64,
	CPUFuzzMatchFactor int,
	RAMFuzzMatchFactor float64,
	roundUp bool,
) []Product {

	var (
		l          []Product
		minC, maxC int
		minR, maxR float64
	)

	index, err := extractEC2Product("data/ec2_offer_" + region + ".json")
	if err != nil {
		log.Fatal(err)
	}

	item := reqInstance{
		CPU: reqCPU,
		RAM: reqRAM,
	}

	// r := make(map[string]Product)
	for _, product := range index.Products {
		if product.Attributes.Vcpu == item.CPU {
			l = append(l, product)
		}
	}
	if len(l) < 1 {
		switch roundUp {
		case false:
			minC, maxC = item.CPU-CPUFuzzMatchFactor, item.CPU+CPUFuzzMatchFactor
			minR, maxR = item.RAM-RAMFuzzMatchFactor, item.RAM+RAMFuzzMatchFactor

		default:
			minC, maxC = item.CPU, item.CPU+CPUFuzzMatchFactor
			minR, maxR = item.RAM, item.RAM+RAMFuzzMatchFactor

		}
		for _, product := range index.Products {
			if (minC <= product.Attributes.Vcpu && product.Attributes.Vcpu <= maxC) &&
				(minR <= product.Attributes.Memory.value && product.Attributes.Memory.value <= maxR) {
				l = append(l, product)
			}
		}
	}
	return l
}
