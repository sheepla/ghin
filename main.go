package main

import (
	"fmt"
	"log"

	"github.com/sheepla/ghin/gh"
)

func main() {
	param := gh.NewSearchParam("vim")
	result, err := gh.Search(param)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
