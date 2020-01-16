package ec2

import (
	"log"
)

// InstancePrefixList : comment
type InstancePrefixList struct {
	InstanceClassPrefix []string
}

// PreferenceAttr : comment
type PreferenceAttr struct {
	InstanceClassPrefix string
	Burstable           bool
	CurrentGen          bool
	Include             InstancePrefixList
	Exclude             InstancePrefixList
}

// SimpleSearchReq : comment
type SimpleSearchReq struct {
	Operation   string
	SearchType  string
	Region      string
	CPU         int
	Memory      RAM
	CPUFF       int
	RAMFF       float64
	UpsizeOnly  bool
	Preferences PreferenceAttr
}

// ResultInstance : comment
type ResultInstance struct {
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

func makeResultInstance(ec2attr Ec2Attributes) ResultInstance {
	return ResultInstance{
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

// SimpleSearch : comment
func SimpleSearch(request SimpleSearchReq) []map[string]ResultInstance {
	var (
		l          []ResultInstance
		minC, maxC int
		minR, maxR float64
	)

	index, err := extractEC2Product("data/ec2_offer_" + request.Region + ".json")
	if err != nil {
		log.Fatal(err)
	}

	// find perfect matches
	for _, product := range index.Products {
		if (product.Attributes.Vcpu == request.CPU) &&
			(product.Attributes.Memory.Value == request.Memory.Value) {
			instance := makeResultInstance(product.Attributes)
			l = append(l, instance)
		}
	}

	// if not perfect matches found, then fuzzy match
	if len(l) < 1 {
		maxC = request.CPU + request.CPUFF
		maxR = request.Memory.Value + request.RAMFF
		switch request.UpsizeOnly {
		case false:
			minC = request.CPU - request.CPUFF
			minR = request.Memory.Value - request.RAMFF
		default:
			minC, minR = request.CPU, request.Memory.Value
		}
		for _, product := range index.Products {
			if (product.Attributes.Vcpu <= maxC && product.Attributes.Vcpu >= minC) &&
				(product.Attributes.Memory.Value <= maxR && product.Attributes.Memory.Value >= minR) {
				instance := makeResultInstance(product.Attributes)
				l = append(l, instance)
			}
		}
	}

	var ulist []map[string]ResultInstance
	keys := make(map[string]bool)
	for _, instance := range l {
		if _, ok := keys[instance.InstanceType]; !ok {
			t := make(map[string]ResultInstance)
			keys[instance.InstanceType] = true
			t[instance.InstanceType] = instance
			ulist = append(ulist, t)
		}
	}

	return ulist

}
