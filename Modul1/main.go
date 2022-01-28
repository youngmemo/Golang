package main

import (
	"example.com/oppgaver/MyQuotes"
	"example.com/oppgaver/RiverCrossing"
	"fmt"
)

func main() {
	fmt.Println(MyQuotes.PrintGo())
	fmt.Println(MyQuotes.PrintGlass())
	fmt.Println(MyQuotes.PrintHello())
	fmt.Println(MyQuotes.PrintOpt())

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
