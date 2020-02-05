package ec2

import (
	"fmt"
	"log"
	"strings"
)

// PreferenceAttr : comment
type PreferenceAttr struct {
	IncludeBurstable        bool
	CurrentGenOnly          bool
	ExcludeInstanceFamilies []string
}

// SimpleSearchReq : comment
type SimpleSearchReq struct {
	WorkloadName string
	Region       string
	CPU          int
	Memory       RAM
	CPUFF        int
	RAMFF        float64
	UpsizeOnly   bool
	Preferences  PreferenceAttr
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

type ResultSet []ResultInstance

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
	f := fmt.Sprintf("%s/%s.json", dataPath, request.Region)
	index, err := extractEC2Product(f)
	if err != nil {
		log.Fatal(err)
	}

	// find perfect matches
	for _, product := range index.Products {
		if (product.Attributes.Vcpu == request.CPU) &&
			(product.Attributes.Memory.Value == request.Memory.Value) {
			if matchedPreferences(request, product.Attributes) {
				l = append(l, makeResultInstance(product.Attributes))
			}
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
				if matchedPreferences(request, product.Attributes) {
					l = append(l, makeResultInstance(product.Attributes))
				}
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

func matchedPreferences(req SimpleSearchReq, ec2Attr Ec2Attributes) bool {
	l := []bool{
		matchGenerationPref(req.Preferences.CurrentGenOnly, ec2Attr.CurrentGeneration),
		matchBurstablePref(req.Preferences.IncludeBurstable, ec2Attr.InstanceType),
		matchInstanceExclusionPref(req.Preferences.ExcludeInstanceFamilies, ec2Attr.InstanceType),
	}

	for _, b := range l {
		if !b {
			return false
		}
	}
	return true
}

func matchGenerationPref(desiredPref bool, recievedVal string) bool {
	if desiredPref {
		var recBool bool
		switch strings.ToLower(recievedVal) {
		case "yes":
			recBool = true
		case "no":
			recBool = false
		default:
			log.Fatal("Unknown Value in isCurrentGen(actual). Should be either 'Yes' or 'No'.")
		}
		return desiredPref == recBool
	}
	// if desiredpref was false, then no check is required
	// hence return true, as match
	return true
}

func matchBurstablePref(desiredPref bool, instanceType string) bool {
	// if desiredPref is set to false, i.e. exlcude burstable instances
	// check if instance name starts with one of the elements of bFamPrefix
	// if it does, then it is a burstable instance, and overall should return false
	if !desiredPref {
		var isBurstable bool
		bFamPrefix := []string{"t"}
		for _, prefix := range bFamPrefix {
			if strings.HasPrefix(instanceType, prefix) {
				isBurstable = true
				break
			}
		}
		// this will return false as isBurstable is true
		// and desiredPref is false
		return desiredPref == isBurstable
	}
	return true
}

func matchInstanceExclusionPref(list []string, instanceType string) bool {
	if len(list) > 0 {
		instancePrefix := strings.Split(instanceType, ".")[0]
		for _, prefix := range list {
			if strings.EqualFold(prefix, instancePrefix) {
				return false
			}
		}
	}
	return true
}
