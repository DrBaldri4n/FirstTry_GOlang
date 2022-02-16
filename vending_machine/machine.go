package main

/*
import (
	"errors"
	"fmt"
)

//calculation of the machine content with error handling
func mainMachine(remove string) error {
	switch remove {
	case "water":
		if waterQuantity == 0 {
			return errors.New("water is empty")
		} else {
			waterQuantity -= 1
			moneyTransfer(0.80)
		}
	case "cola":
		if colaQuantity == 0 {
			return errors.New("cola is empty")
		} else {
			colaQuantity -= 1
			moneyTransfer(1.50)
		}
	case "soda":
		if sodaQuantity == 0 {
			return errors.New("soda is empty")
		} else {
			sodaQuantity -= 1
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
	costsInt := round(costs)
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
*/
