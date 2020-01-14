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
	// case "findPrice":
	// 	src, dest, region := args[1], args[2], args[3]
	// 	err := ec2.GetEC2Prices(
	// 		src,
	// 		dest,
	// 		region,
	// 	)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// case "findMatch":
	// 	if len(args) != 5 {
	// 		fmt.Println("")
	// 		fmt.Printf("Error: Required 4 arguments. Provided: %d\n", len(args))
	// 		fmt.Println("--------")
	// 		fmt.Printf("Usage: go run main.go findMatch [SRC_CSV_PATH] [REGION] [CPUFuzzFactor] [RAMFuzzFactor]\n")
	// 		log.Fatal()
	// 	}
	// 	src, region := args[1], args[2]
	// 	CPUFuzzFactor, err := strconv.Atoi(args[3])
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	RAMFuzzFactor, err := strconv.ParseFloat(args[4], 64)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	result, err := ec2.GetEC2Match(src, region, CPUFuzzFactor, RAMFuzzFactor)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(result)
	case "byCPU":
		if len(args) != 7 {
			fmt.Println("")
			fmt.Printf("Error: Required 4 arguments. Provided: %d\n", len(args))
			fmt.Println("--------")
			fmt.Printf("Usage: go run main.go byCPU [REGION] [CPU int] [RAM float64] [CPUFuzzFactor int] [RAMFuzzFactor float64] [roundUp bool]\n")
			log.Fatal()
		}

		region := args[1]

		reqCPU, err := strconv.Atoi(args[2])
		if err != nil {
			log.Fatal(err)
		}

		reqRAM, err := strconv.ParseFloat(args[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		cff, err := strconv.Atoi(args[4])
		if err != nil {
			log.Fatal(err)
		}

		rff, err := strconv.ParseFloat(args[5], 64)
		if err != nil {
			log.Fatal(err)
		}

		roundUp, err := strconv.ParseBool(args[6])
		if err != nil {
			log.Fatal(err)
		}

		result := ec2.ByCPUAndRAM(region, reqCPU, reqRAM, cff, rff, roundUp)
		if err != nil {
			log.Fatal(err)
		}

		for _, v := range result {
			fmt.Println(v.Attributes.InstanceType)
		}

		fmt.Println(len(result))
	default:
		log.Fatal("Case not found")
	}

}
