package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/rmaidveo/go-life"
)

func main() {
	outDelay := flag.Duration("delay", 100*time.Millisecond, "delay between frames")
	flag.Parse()

	text, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("unable to read the field: %s", err)
	}

	field, err := life.ParseField(string(text))
	if err != nil {
		log.Fatalf("unable to parse the field: %s", err)
	}

	for {
		fmt.Print(field)
		time.Sleep(*outDelay)

		field = field.NextField()
	}
}
