package ec2

type offer struct {
	OfferCode                  string `json:"offerCode"`
	VersionIndexURL            string `json:"versionIndexUrl"`
	CurrentVersionURL          string `json:"currentVersionUrl"`
	CurrentRegionIndexURL      string `json:"currentRegionIndexUrl"`
	SavingsPlanVersionIndexURL string `json:"savingsPlanVersionIndexUrl,omitempty"`
	CurrentSavingsPlanIndexURL string `json:"currentSavingsPlanIndexUrl,omitempty"`
}

type offerIndex struct {
	FormatVersion   string           `json:"formatVersion"`
	Disclaimer      string           `json:"disclaimer"`
	PublicationDate string           `json:"publicationDate"`
	Offers          map[string]offer `json:"offers"`
}

type region struct {
	RegionCode        string `json:"regionCode"`
	CurrentVersionURL string `json:"currentVersionUrl"`
}

type regionIndex struct {
	FormatVersion   string            `json:"formatVersion"`
	Disclaimer      string            `json:"disclaimer"`
	PublicationDate string            `json:"publicationDate"`
	Regions         map[string]region `json:"regions"`
}

type ec2Attributes struct {
	Servicecode                 string `json:"servicecode"`
	Location                    string `json:"location"`
	LocationType                string `json:"locationType"`
	InstanceType                string `json:"instanceType"`
	CurrentGeneration           string `json:"currentGeneration"`
	InstanceFamily              string `json:"instanceFamily"`
	Vcpu                        string `json:"vcpu"`
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
	value float64
	unit  string
}

type ec2 struct {
	Sku           string        `json:"sku"`
	ProductFamily string        `json:"productFamily"`
	Attributes    ec2Attributes `json:"attributes"`
}

type ec2ProdIndex struct {
	FormatVersion   string         `json:"formatVersion"`
	Disclaimer      string         `json:"disclaimer"`
	PublicationDate string         `json:"publicationDate"`
	Products        map[string]ec2 `json:"products"`
	Terms           terms          `json:"terms"`
}

type terms struct {
	OnDemand map[string]map[string]offerTerm `json:"OnDemand"`
	Reserved map[string]map[string]offerTerm `json:"Reserved,omitempty"`
}

type offerTerm struct {
	OfferTermCode   string                     `json:"offerTermCode"`
	Sku             string                     `json:"sku"`
	EffectiveDate   string                     `json:"effectiveDate"`
	PriceDimensions map[string]priceDimensions `json:"priceDimensions"`
	TermAttributes  termAttributes             `json:"termAttributes"`
}
type pricePerUnit struct {
	USD string
}

type priceDimensions struct {
	RateCode     string       `json:"rateCode"`
	Description  string       `json:"description"`
	BeginRange   string       `json:"beginRange"`
	Unit         string       `json:"unit"`
	PricePerUnit pricePerUnit `json:"pricePerUnit"`
	AppliesTo    []string     `json:"appliesTo"`
}

type termAttributes struct {
	LeaseContractLength string `json:"LeaseContractLength"`
	OfferingClass       string `json:"OfferingClass"`
	PurchaseOption      string `json:"PurchaseOption"`
}
