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
type resultInstance struct {
	InstanceType                string
	CurrentGeneration           string
	InstanceFamily              string
	Vcpu                        int
	Memory                      RAM
	Storage                     string
	NetworkPerformance          string
	ProcessorArchitecture       string
	Tenancy                     string
	DedicatedEbsThroughput      string
	EnhancedNetworkingSupported string
}

func makeResultInstance(ec2attr Ec2Attributes) resultInstance {
	return resultInstance{
		InstanceType:                ec2attr.InstanceType,
		CurrentGeneration:           ec2attr.CurrentGeneration,
		InstanceFamily:              ec2attr.InstanceFamily,
		Vcpu:                        ec2attr.Vcpu,
		Memory:                      ec2attr.Memory,
		Storage:                     ec2attr.Storage,
		Tenancy:                     ec2attr.Tenancy,
		NetworkPerformance:          ec2attr.NetworkPerformance,
		ProcessorArchitecture:       ec2attr.ProcessorArchitecture,
		DedicatedEbsThroughput:      ec2attr.DedicatedEbsThroughput,
		EnhancedNetworkingSupported: ec2attr.EnhancedNetworkingSupported,
	}
}

// ByCPUAndRAM : comment
func ByCPUAndRAM(
	region string,
	reqCPU int,
	reqRAM float64,
	CPUFuzzMatchFactor int,
	RAMFuzzMatchFactor float64,
	roundUp bool,
	unique bool,
) []map[string]resultInstance {

	var (
		l          []resultInstance
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

	// find perfect matches
	for _, product := range index.Products {
		if (product.Attributes.Vcpu == item.CPU) &&
			(product.Attributes.Memory.value == item.RAM) {
			instance := makeResultInstance(product.Attributes)
			l = append(l, instance)
		}
	}

	// if not perfect matches found, then fuzzy match
	if len(l) < 1 {
		maxC = item.CPU + CPUFuzzMatchFactor
		maxR = item.RAM + RAMFuzzMatchFactor
		switch roundUp {
		case false:
			minC = item.CPU - CPUFuzzMatchFactor
			minR = item.RAM - RAMFuzzMatchFactor
		case true:
			minC, minR = item.CPU, item.RAM
		default:
			minC, minR = item.CPU, item.RAM
		}
		for _, product := range index.Products {
			if (product.Attributes.Vcpu <= maxC && product.Attributes.Vcpu >= minC) &&
				(product.Attributes.Memory.value <= maxR && product.Attributes.Memory.value >= minR) {
				instance := makeResultInstance(product.Attributes)
				l = append(l, instance)
			}
		}
	}

	var ulist []map[string]resultInstance
	keys := make(map[string]bool)
	for _, instance := range l {
		if _, ok := keys[instance.InstanceType]; !ok {
			t := make(map[string]resultInstance)
			keys[instance.InstanceType] = true
			t[instance.InstanceType] = instance
			ulist = append(ulist, t)
		}
	}

	return ulist

}
