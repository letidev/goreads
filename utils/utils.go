package utils

import "fmt"

// Prints the error message if not nil
func PrintIfErr(err error, message string) {
	if err != nil {
		fmt.Println(err)
		fmt.Println("==============")
		fmt.Println(message)
	}
}

// Panics if error not nil
func PanicIfEff(err error, message string) {
	if err != nil {
		fmt.Println(err)
		fmt.Println("==============")
		panic(message)
	}
}
