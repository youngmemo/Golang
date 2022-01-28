package main

import (
	"example.com/oppgaver/RiverCrossing"
	"fmt"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Opt())

	fmt.Println(RiverCrossing.ViewState())
	RiverCrossing.PutInFox()
	fmt.Println(RiverCrossing.ViewState())
	fmt.Println("")

	RiverCrossing.PutInHS()
	fmt.Println(RiverCrossing.ViewState())
	fmt.Println("")

	RiverCrossing.PutInCorn()
	fmt.Println(RiverCrossing.ViewState())
	fmt.Println("")

	RiverCrossing.PutInMonkey()
	fmt.Println(RiverCrossing.ViewState())
	fmt.Println("")

	RiverCrossing.PutInWhat("HS KYLLING KORN")
}
