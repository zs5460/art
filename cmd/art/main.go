package main

import (
	"flag"
	"fmt"

	"github.com/wonx2/art"
)

var s string

func main() {
	flag.StringVar(&s, "text", "art - zs5460", "")
	flag.Parse()
	fmt.Println(art.String(s))
}
