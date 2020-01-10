package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mwaqqas/go-ec2-matcher/pkg/ec2"
)

func main() {
	args := os.Args[1:]
	operation := args[0]
	switch operation {
	case "findPrice":
		src, dest, region := args[1], args[2], args[3]
		err := ec2.GetEC2Prices(
			src,
			dest,
			region,
		)
		if err != nil {
			log.Fatal(err)
		}
	case "findMatch":
		src, region := args[1], args[2]
		result, err := ec2.GetEC2Match(src, region)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	default:
		log.Fatal("Case not found")
	}

}
