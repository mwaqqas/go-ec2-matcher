package ec2

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
//
// 
// )

// func downloadFile(url string, file string, force bool) {
// 	// Check if File needs to be downloaded or not
// 	var isExist bool
// 	if _, err := os.Stat(file); err == nil {
// 		isExist = true
// 	}
// 	if force == false && isExist == true {
// 		fmt.Printf("%v already exists and force download is set to false.\n", file)
// 		return
// 	}

// 	// Create file
// 	f, err := os.Create(file)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer f.Close()

// 	// Download data
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	// Copy downloaded data to file
// 	_, err = io.Copy(f, resp.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	return
// }

// func downloadEC2RegionOffers(file string) (err error) {
// 	// Read file
// 	f, err := ioutil.ReadFile(file)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	// parse the offer index
// 	var rIndex regionIndex
// 	err = json.Unmarshal([]byte(f), &rIndex)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	for _, v := range rIndex.Regions {
// 		url := baseURL + v.CurrentVersionURL
// 		path := dataDirPath + "/ec2_offer_" + v.RegionCode + ".json"
// 		fmt.Printf("Downloading EC2 Offer file for Region: %v\n", v.RegionCode)
// 		downloadFile(url, path, false)
// 	}
// 	return
// }
