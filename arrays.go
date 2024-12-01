package main

import "fmt"

type Product struct{
	title string
	id int
	price float64
}

func main(){
	//Practice
	hobbies := [3]string{"coding", "sports", "movies"}

	// 1. 
	fmt.Print(hobbies)

	// 2.
	fmt.Print(hobbies[0])
	fmt.Print(hobbies[1:])

	// 3.
	mainHobbies := hobbies[:2]
	fmt.Print(mainHobbies)
	// fmt.Print(hobbies[0:2])

	// 4.
	mainHobbies = mainHobbies[1:3]
	fmt.Print(mainHobbies)

	// 5.
	var goals []string = []string{"tm", "cto"}
	fmt.Print(goals)

	// 6.
	goals[1] = "ceo"
	goals = append(goals, "cto" )

	fmt.Print(goals)
	
	// 7.
	products := []Product{Product{"Iphone", 1, 100}, Product{"samsung", 2, 150}}
	products = append(products, Product{"Xiaomi", 3, 50})
	fmt.Print(products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first array that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.