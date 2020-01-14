package ec2

import (
	"log"
)

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
