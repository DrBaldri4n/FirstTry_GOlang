package main

import (
	"errors"
	"fmt"
)

/*
var waterQuantity int // sollte ich int64 nutzten? //wie ist das mit Globalen Var?
var colaQuantity int  //struct
var sodaQuantity int
*/

type fillMachine1 struct {
	waterQuantity int
	colaQuantity  int
	sodaQuantity  int
}

/*
func fillMachine() {
	waterQuantity = 1
	colaQuantity = 1
	sodaQuantity = 3
}
*/

func main() {
	//init
	var userInput int
	var userString string
	exit := 1

	fillMachine := fillMachine1{3, 3, 3}

	//Main func
	for exit != 0 {
		fmt.Println("\nchoose a drink with nummbers")
		if fillMachine.waterQuantity != 0 {
			fmt.Println("1 - Water \t 0,80 €")
		}
		if fillMachine.colaQuantity != 0 {
			fmt.Println("2 - Cola \t 1,50 €")
		}
		if fillMachine.sodaQuantity != 0 {
			fmt.Println("3 - soda \t 1,20 €")
		}
		fmt.Println("4 - exit machine")

		//userInput
		fmt.Scan(&userInput)

		//transfer userInput
		switch userInput {
		case 1:
			userString = "water"
		case 2:
			userString = "cola"
		case 3:
			userString = "soda"
		case 4:
			exit = 0
			continue
		}
		if e := mainMachine(userString, fillMachine); e != nil {
			fmt.Println("Error: ", e)
		}

		contentsPrint("Content: ", fillMachine)
	}
	fmt.Println("\n\nmachine is finished")
	contentsPrint("Rest: ", fillMachine)
}

//calculation of the machine content with error handling
func mainMachine(remove string, fM1 fillMachine1) error {
	switch remove {
	case "water":
		fmt.Println("Wasser: ", fM1.waterQuantity)
		if fM1.waterQuantity == 0 {
			return errors.New("water is empty")
		} else {
			fM1.waterQuantity -= 1
			fmt.Println("Wasser danach: ", fM1.waterQuantity)
			moneyTransfer(0.80)
		}
	case "cola":
		if fM1.colaQuantity == 0 {
			return errors.New("cola is empty")
		} else {
			fM1.colaQuantity -= 1
			moneyTransfer(1.50)
		}
	case "soda":
		if fM1.sodaQuantity == 0 {
			return errors.New("soda is empty")
		} else {
			fM1.sodaQuantity -= 1
			moneyTransfer(1.20)
		}
	default:
		return errors.New("error in the remove string transfer")
	}
	fmt.Println("take your ", remove, " drink :)")
	fmt.Println()
	return nil
}

func moneyTransfer(costs float32) { //sollte ich hier auf performance achten bei größeren Projekten?
	var input float32
	costsInt := round(costs) //math.round
	for costs > 0 {
		//input
		fmt.Println("throw in 0,10 | 0,20 | 0,50 | 1,00 | 2,00  €")
		fmt.Println("remaining: ", costsInt, "CENT")
		fmt.Scan(&input)
		inputInt := round(input)

		switch inputInt {
		case 10:
			costsInt -= inputInt
		case 20:
			costsInt -= inputInt
		case 50:
			costsInt -= inputInt
		case 100:
			costsInt -= inputInt
		case 200:
			costsInt -= inputInt
		default:
			fmt.Println("only EURO coins possible")
		}

		//output
		if costsInt < 0 {
			fmt.Println("take your money")
			costsInt *= (-1)
			for costsInt >= 100 {
				costsInt -= 100
				fmt.Println("1,00 €")
			}
			for costsInt >= 50 {
				costsInt -= 50
				fmt.Println("0,50 €")
			}
			for costsInt >= 20 {
				costsInt -= 20
				fmt.Println("0,20 €")
			}
			for costsInt >= 10 {
				costsInt -= 10
				fmt.Println("0,10 €")
			}
		}

		if costsInt == 0 {
			break
		}
	}
}

//Output
func contentsPrint(name string, fM1 fillMachine1) {
	fmt.Println(name)
	fmt.Printf("Water: %d\nCola: %d\nSoda: %d", fM1.waterQuantity, fM1.colaQuantity, fM1.sodaQuantity)
}

//round
func round(val float32) int {
	val1 := val * 100
	if val1 < 0 {
		return int(val1)
	}
	return int(val1)
}
