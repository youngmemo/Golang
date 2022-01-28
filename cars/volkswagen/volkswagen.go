package volkswagen

import "fmt"

var make = "Memo"
var model = "Polo"
var year = 2011
var wheels = 4
var models []int

func AddModels() {
	for i := 0; i < 50; i++ {
		models = append(models, i)
	}
}

func AddOneModel(YearOnModel int) {
	models = append(models, YearOnModel)
}

func FindModel(WantModel int) {
	for i := 0; i < len(models); i++ {
		if i == WantModel {
			fmt.Printf("We have the model! \n")
		}
	}
}

func GetModels() {
	fmt.Println(models)
}
