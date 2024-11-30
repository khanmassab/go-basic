package main

import "fmt"

func main(){

	var stringArray [4]string = [4]string{"element1", "element2"}
	var prices [4] float64 = [4] float64{64.0, 43.78, 453, 654}

	stringArray[2] = "element3"
	fmt.Println(prices[0:2]) // from index 0 and before index 2, i.e, 0 and 1
	fmt.Println(stringArray[1])
}