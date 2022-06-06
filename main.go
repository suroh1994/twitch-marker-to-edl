package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("incorrect usage: program csvfile")
		return
	}
	csvFile := os.Args[1]
	Export, err := ImportCSV(csvFile)
	fmt.Printf("%v\n%v", err, Export)
}
