package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/mwaqqas/go-ec2-matcher/pkg/ec2"
)

func main() {
	args := os.Args[1:]
	operation := args[0]
	switch operation {
	case "SimpleSearch":
		if len(args) != 2 {
			fmt.Println("")
			fmt.Printf("Error: Required 2 arguments. Provided: %d\n", len(args))
			log.Fatal()
		}

		b := []byte(args[1])
		var req ec2.SimpleSearchReq
		err := json.Unmarshal(b, &req)
		if err != nil {
			log.Fatal(err)
		}
		result := ec2.SimpleSearch(req)
		out, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(out))

	default:
		log.Fatal("Case not found")
	}

}
