package pointers

import "fmt"

func main(){
	age := 32

	agePointer := &age
	fmt.Println("Age: ", *agePointer)

	setAdultYears(agePointer)
	fmt.Println("Adult Years: ", getAge(agePointer))
}

func getAge(age *int) int{
	return *age;
}

func setAdultYears(age *int){
	//Deference a pointer i.e. &age
	*age = *age - 18
}