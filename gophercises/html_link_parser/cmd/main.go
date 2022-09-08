package main

import (
	"fmt"
	"strings"

	"htmllinkparser/internal/examples"
	"htmllinkparser/internal/parser"
)

func main() {
	r := strings.NewReader(examples.Ex1)

	links, err := parser.Parse(r)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Printf("%+v\n\n ---------- Main finished ----------\n\n", links)
}
