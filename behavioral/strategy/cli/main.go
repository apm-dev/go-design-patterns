package main

import (
	"flag"
	"github.com/apm-dev/go-design-patterns/behavioral/strategy/shapes"
	"log"
	"os"
)

var output = flag.String("output", "text", "where to write output, 'console' or 'image' file")

func main() {
	flag.Parse()
	activeStrategy, err := shapes.NewPrinter(*output)
	if err != nil {
		log.Fatal(err)
	}
	switch *output {
	case shapes.TextStrategy:
		activeStrategy.SetWriter(os.Stdout)
	case shapes.ImageStrategy:
		w, err := os.Create("/tmp/image.jpg")
		if err != nil {
			log.Fatal("error opening image")
		}
		defer w.Close()
		activeStrategy.SetWriter(w)
	}
	err = activeStrategy.Print()
	if err != nil {
		log.Fatal(err)
	}
}
