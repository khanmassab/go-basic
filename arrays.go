package main

import "fmt"

func main(){
	prices := []float64{1, 2}

	newPrices := append(prices, 43)

	fmt.Print(newPrices)
}

// func main(){

// 	var stringArray [4]string = [4]string{"element1", "element2"}
// 	var prices [4] float64 = [4] float64{64.0, 43.78, 453, 654}

// 	var fromSecondToTheLastIndex = prices[1:]

// 	fmt.Print(fromSecondToTheLastIndex)
// 	stringArray[2] = "element3"
// 	fmt.Println(prices[0:2]) // from index 0 and before index 2, i.e, 0 and 1
// 	fmt.Println(stringArray[1])
// 	fmt.Print(len(fromSecondToTheLastIndex), cap(fromSecondToTheLastIndex))
// }