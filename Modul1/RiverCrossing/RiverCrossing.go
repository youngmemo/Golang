package RiverCrossing

import "fmt"

var state = "Nothing is in the boat"

func ViewState() string {
	return state
}

func PutInFox() {
	state = "Fox is in the boat"
}

func PutInWhat(what string) {
	fmt.Printf("\\___%s___/", what)
}

func PutInHS() {
	state = "HS is in the boat"
}

func PutInCorn() {
	state = "Corn is in the boat"
}

func PutInMonkey() {
	state = "Monkey is in the boat"
}
