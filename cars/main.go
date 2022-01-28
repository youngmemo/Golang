package main

import (
	"fmt"
	"memo.com/cars/volkswagen"
)

func main() {
	fmt.Println("I love cars! One of the cars I have is in my garage, I love it!")

	volkswagen.AddModels()
	volkswagen.AddOneModel(1234)
	volkswagen.FindModel(6)
	volkswagen.GetModels()
}
