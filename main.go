package main

import (
	"fmt"
	"log"

	"github.com/bgildson/gomock-example/client/finalspace0"
)

func main() {
	fsClient := finalspace0.New()

	quotes, err := fsClient.GetQuotes()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", quotes)
}
