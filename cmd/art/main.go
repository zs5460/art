package main

import (
	"flag"
	"fmt"

	"github.com/zs5460/art"
)

var s string

func main() {
	flag.StringVar(&s, "text", "art - zs5460", "")
	flag.Parse()
	fmt.Println(art.String(s))
}
