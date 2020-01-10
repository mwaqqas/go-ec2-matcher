package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

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
		if len(args) != 5 {
			fmt.Println("")
			fmt.Printf("Error: Required 4 arguments. Provided: %d\n", len(args))
			fmt.Println("--------")
			fmt.Printf("Usage: go run main.go findMatch [SRC_CSV_PATH] [REGION] [CPUFuzzFactor] [RAMFuzzFactor]\n")
			log.Fatal()
		}
		src, region := args[1], args[2]
		CPUFuzzFactor, err := strconv.Atoi(args[3])
		if err != nil {
			log.Fatal(err)
		}
		RAMFuzzFactor, err := strconv.ParseFloat(args[4], 64)
		if err != nil {
			log.Fatal(err)
		}
		result, err := ec2.GetEC2Match(src, region, CPUFuzzFactor, RAMFuzzFactor)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	default:
		log.Fatal("Case not found")
	}

}
