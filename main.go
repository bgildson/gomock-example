package main

import (
	"fmt"
	"log"
	"net/http"

	// "github.com/bgildson/gomock-example/client/finalspace0"
	// finalspace "github.com/bgildson/gomock-example/client/finalspace1"
	// finalspace "github.com/bgildson/gomock-example/client/finalspace2"
	finalspace "github.com/bgildson/gomock-example/client/finalspace3"
)

func main() {
	// fsClient := finalspace0.New()
	fsClient := finalspace.New(http.DefaultClient, "https://finalspaceapi.com")

	quotes, err := fsClient.GetQuotes()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", quotes)
}
