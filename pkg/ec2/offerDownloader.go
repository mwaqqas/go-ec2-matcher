package ec2

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type OfferRequest struct {
	Regions    []string `json:"regions,omitempty"`
	Force      bool     `json:"force,omitempty"`
	AllRegions bool     `json:"allRegions,omitempty"`
}

func GetEC2OfferFiles(request OfferRequest) {
	// download global offer index
	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		err = os.MkdirAll(dataPath, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
	err := DownloadFile(offerIndexURL, offerIndexPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Read Global Offer Index and get region index path for EC2
	regionIndexURI, err := GetEC2RegionIndexURL(offerIndexPath)
	if err != nil {
		log.Fatal(err)
	}

	// download region index for EC2
	regionIndexURL := baseURL + regionIndexURI
	err = DownloadFile(regionIndexURL, regionIndexPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	var regions []string
	switch request.AllRegions {
	case true:
		regions, err = GetListOfAllRegions(regionIndexPath)
		if err != nil {
			log.Fatal(err)
		}
	default:
		regions = request.Regions
	}
	var count int
	fmt.Printf("Region Offers to Download: %d\n", len(regions))
	for _, region := range regions {
		url, err := GetURLForRegionEC2Offer(regionIndexPath, region)
		if err != nil {
			log.Fatal(err)
		}
		fileName := fmt.Sprintf("%s/%s.json", dataPath, region)
		switch request.Force {
		case true:
			err = DownloadFile(url, fileName)
			if err != nil {
				log.Fatal(err)
			}
		default:
			if _, err := os.Stat(fileName); os.IsNotExist(err) {
				err = DownloadFile(url, fileName)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		count++
		fmt.Printf("Downloaded: %d of %d\n", count, len(regions))
	}
}

func DownloadFile(url string, file string) (err error) {
	// Create file
	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer f.Close()

	// Download data
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer resp.Body.Close()

	// Copy downloaded data to file
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return
}

func GetEC2RegionIndexURL(file string) (url string, err error) {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
		return
	}
	var index OfferIndex
	err = json.Unmarshal([]byte(f), &index)
	for k, v := range index.Offers {
		if k == "AmazonEC2" {
			url = v.CurrentRegionIndexURL
			return
		}
	}
	return
}

func GetURLForRegionEC2Offer(file string, region string) (url string, err error) {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
		return
	}
	var index RegionIndex
	err = json.Unmarshal([]byte(f), &index)
	for _, r := range index.Regions {
		if r.RegionCode == region {
			url = baseURL + r.CurrentVersionURL
			return
		}
	}
	return
}

func GetListOfAllRegions(file string) (regions []string, err error) {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
		return
	}
	var index RegionIndex
	err = json.Unmarshal([]byte(f), &index)
	if err != nil {
		log.Fatal(err)
	}
	for k := range index.Regions {
		regions = append(regions, k)
	}
	return
}
