package main

import (
	"fmt"
	"strings"

	"htmllinkparser/pkg/examples"
	"htmllinkparser/pkg/parser"
)

func main() {
	r := strings.NewReader(examples.Ex3)

	links, err := link.Parse(r)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Printf("%+v\n", links)
}
