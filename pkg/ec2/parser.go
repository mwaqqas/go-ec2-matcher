package ec2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func extractEC2Offer(file string) (o offer, err error) {
	// Read file
	f, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	// parse the offer index
	var oIndex offerIndex
	err = json.Unmarshal([]byte(f), &oIndex)
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, v := range oIndex.Offers {
		if k == "AmazonEC2" {
			return v, err
		}
	}
	return
}

func extractEC2Product(file string) (index ec2ProdIndex, err error) {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Read file Error")
		return
	}
	err = json.Unmarshal([]byte(f), &index)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Unmarshall Error, extractEC2Product")
		return
	}
	return
}
