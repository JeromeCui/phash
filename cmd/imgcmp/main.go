package main

import (
	"fmt"
	"os"

	"github.com/JeromeCui/phash"
)

func main() {
	if len(os.Args) <= 2 {
		fmt.Printf("need source file and target file\n")
		return
	}

	h1, err := phash.ImageHashDCT(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	h2, err := phash.ImageHashDCT(os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("h1: ", h1)
	fmt.Println("h2: ", h2)
	fmt.Println("distance: ", phash.HammingDistance(h1, h2))
}
