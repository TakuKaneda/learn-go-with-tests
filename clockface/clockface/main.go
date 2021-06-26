package main

import (
	"os"
	"time"

	"learn-go-with-tests/clockface"
	// "github.com/quii/learn-go-with-tests/math/clockface" // REPLACE THIS!
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
