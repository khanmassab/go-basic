package main

import "fmt"

func main(){

	var stringArray [4]string = [4]string{"element1", "element2"}

	stringArray[2] = "element3"

	fmt.Print(stringArray[1])
}