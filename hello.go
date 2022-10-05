package main

import (
	"fmt"
	"helloMyWorld/playingWithPackages"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())

	sum := playingWithPackages.Add(1, 5)
	fmt.Println(sum)
}
