package main

import (
	"os"
	"time"

	"github.com/ammpedro/learn-go-with-tests/mathclock"
)

func main() {
	t := time.Now()
	mathclock.SVGWriter(os.Stdout, t)
}
