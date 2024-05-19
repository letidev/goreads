package utils

import "fmt"

func PrintIfErr(err error, message string) {
	if err != nil {
		fmt.Println(err)
		fmt.Println("==============")
		fmt.Println(message)
	}
}

func PanicIfEff(err error, message string) {
	if err != nil {
		fmt.Println(err)
		fmt.Println("==============")
		panic(message)
	}
}
