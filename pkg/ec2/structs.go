package ec2

type Offer struct {
	OfferCode                  string `json:"offerCode"`
	VersionIndexURL            string `json:"versionIndexUrl"`
	CurrentVersionURL          string `json:"currentVersionUrl"`
	CurrentRegionIndexURL      string `json:"currentRegionIndexUrl"`
	SavingsPlanVersionIndexURL string `json:"savingsPlanVersionIndexUrl,omitempty"`
	CurrentSavingsPlanIndexURL string `json:"currentSavingsPlanIndexUrl,omitempty"`
}

type OfferIndex struct {
	FormatVersion   string           `json:"formatVersion"`
	Disclaimer      string           `json:"disclaimer"`
	PublicationDate string           `json:"publicationDate"`
	Offers          map[string]Offer `json:"offers"`
}

type Region struct {
	RegionCode        string `json:"regionCode"`
	CurrentVersionURL string `json:"currentVersionUrl"`
}

type RegionIndex struct {
	FormatVersion   string            `json:"formatVersion"`
	Disclaimer      string            `json:"disclaimer"`
	PublicationDate string            `json:"publicationDate"`
	Regions         map[string]Region `json:"regions"`
}

type Ec2Attributes struct {
	Servicecode                 string `json:"servicecode"`
	Location                    string `json:"location"`
	LocationType                string `json:"locationType"`
	InstanceType                string `json:"instanceType"`
	CurrentGeneration           string `json:"currentGeneration"`
	InstanceFamily              string `json:"instanceFamily"`
	Vcpu                        int    `json:"vcpu"`
	PhysicalProcessor           string `json:"physicalProcessor"`
	ClockSpeed                  string `json:"clockSpeed"`
	Memory                      RAM    `json:"memory"`
	Storage                     string `json:"storage"`
	NetworkPerformance          string `json:"networkPerformance"`
	ProcessorArchitecture       string `json:"processorArchitecture"`
	Tenancy                     string `json:"tenancy"`
	OperatingSystem             string `json:"operatingSystem"`
	LicenseModel                string `json:"licenseModel"`
	Usagetype                   string `json:"usagetype"`
	Operation                   string `json:"operation"`
	Capacitystatus              string `json:"capacitystatus"`
	DedicatedEbsThroughput      string `json:"dedicatedEbsThroughput"`
	Ecu                         string `json:"ecu"`
	EnhancedNetworkingSupported string `json:"enhancedNetworkingSupported"`
	Instancesku                 string `json:"instancesku,omitempty"`
	IntelAvxAvailable           string `json:"intelAvxAvailable"`
	IntelAvx2Available          string `json:"intelAvx2Available"`
	IntelTurboAvailable         string `json:"intelTurboAvailable"`
	NormalizationSizeFactor     string `json:"normalizationSizeFactor"`
	PreInstalledSw              string `json:"preInstalledSw"`
	ProcessorFeatures           string `json:"processorFeatures"`
	Servicename                 string `json:"servicename"`
}

// RAM is
type RAM struct {
	Value float64
	Unit  string
}

type Ec2 struct {
	Sku           string        `json:"sku"`
	ProductFamily string        `json:"productFamily"`
	Attributes    Ec2Attributes `json:"attributes"`
}

// Product : comment
type Product struct {
	Sku           string        `json:"sku"`
	ProductFamily string        `json:"productFamily"`
	Attributes    Ec2Attributes `json:"attributes"`
}

type Ec2ProdIndex struct {
	FormatVersion   string             `json:"formatVersion"`
	Disclaimer      string             `json:"disclaimer"`
	PublicationDate string             `json:"publicationDate"`
	Products        map[string]Product `json:"products"`
	// Products        map[string]Ec2 `json:"products"`
	Terms Terms `json:"terms"`
}

type Terms struct {
	OnDemand map[string]map[string]OfferTerm `json:"OnDemand"`
	Reserved map[string]map[string]OfferTerm `json:"Reserved,omitempty"`
}

type OfferTerm struct {
	OfferTermCode   string                     `json:"offerTermCode"`
	Sku             string                     `json:"sku"`
	EffectiveDate   string                     `json:"effectiveDate"`
	PriceDimensions map[string]PriceDimensions `json:"priceDimensions"`
	TermAttributes  TermAttributes             `json:"termAttributes"`
}
type PricePerUnit struct {
	USD string
}

type PriceDimensions struct {
	RateCode     string       `json:"rateCode"`
	Description  string       `json:"description"`
	BeginRange   string       `json:"beginRange"`
	Unit         string       `json:"unit"`
	PricePerUnit PricePerUnit `json:"pricePerUnit"`
	AppliesTo    []string     `json:"appliesTo"`
}

type TermAttributes struct {
	LeaseContractLength string `json:"LeaseContractLength"`
	OfferingClass       string `json:"OfferingClass"`
	PurchaseOption      string `json:"PurchaseOption"`
}
