package main

import "fmt"

type TypeCaster[T any] struct {
	status  bool
	message string
	data    T
}

func main() {
	variableTwo := TypeCaster[any]{status: true, message: "true", data: 8}
	response := processData(variableTwo)
	fmt.Print(response)

}

func processData [TypeCaster any] (response TypeCaster) TypeCaster {
	return response
}
