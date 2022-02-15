package main

import (
	"errors"
	"fmt"
)

func main() {
	//init
	var key int
	var userString string

	//userInput
	fmt.Println("Welche verschluesselung wollen Sie? ")
	fmt.Scan(&key)
	fmt.Println("Geben Sie ihre Nachricht ein")
	fmt.Scan(&userString) //er liest den userString als byteArray ein, stimmt das?

	//Kontrolle des UserInputs
	if _, e := inputControll(userString); e != nil {
		fmt.Println("Error: ", e)
	} else {
		//var transfer
		encryption(key, userString)
	}
}

func encryption(key int, userString string) {

	//init
	alpha := []byte("abcdefghijklmnopqrstuvwxyz")
	decrypt := []byte("")
	output := []byte("")

	//decryped alphaString
	for i := key; i < len(alpha); i++ {
		decrypt = append(decrypt, alpha[i])
	}
	for i := 0; i < key; i++ {
		decrypt = append(decrypt, alpha[i])
	}

	//decryped UserString
	for i := 0; i <= (len(userString) - 1); i++ {
		for letters := 0; letters < len(alpha); letters++ {
			if userString[i] == alpha[letters] {
				output = append(output, decrypt[letters])
				continue
			}
		}
	}

	fmt.Println("decrypted String: ", string(output))
}

func inputControll(userString string) (int, error) {
	for i := 0; i < len(userString); i++ {
		if userString[i] > 96 && userString[i] < 123 {
			continue
		} else {
			return -1, errors.New("nur Eingabe von kleinbuchstaben moeglich")
		}
	}
	return 0, nil
}
