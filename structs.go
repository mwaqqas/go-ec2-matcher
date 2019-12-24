package main

type offer struct {
	OfferCode                  string `json: "offerCode"`
	VersionIndexURL            string `json: "versionIndexUrl"`
	CurrentVersionURL          string `json: "currentVersionUrl"`
	CurrentRegionIndexURL      string `json: "currentRegionIndexUrl"`
	SavingsPlanVersionIndexURL string `json: "savingsPlanVersionIndexUrl,omitempty"`
	CurrentSavingsPlanIndexURL string `json: "currentSavingsPlanIndexUrl,omitempty"`
}

type offerIndex struct {
	FormatVersion   string           `json: "formatVersion"`
	Disclaimer      string           `json: "disclaimer"`
	PublicationDate string           `json: "publicationDate"`
	Offers          map[string]offer `json: "offers"`
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
