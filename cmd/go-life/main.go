package main

import (
	"fmt"
	"log"
	"time"

	"github.com/rmaidveo/go-life"
)

func main() {
	const text = ".....\n.....\n.OOO.\n.....\n....."
	const outDelay = 100 * time.Millisecond

	field, err := life.ParseField(text)
	if err != nil {
		log.Fatalf("unable to parse the field: %s", err)
	}

	for {
		fmt.Print(field)
		time.Sleep(outDelay)

		field = field.NextField()
	}
}
