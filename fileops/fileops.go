package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteBalanceToFile(input float64, fileName string){
	inputText := fmt.Sprint(input)
	os.WriteFile(fileName,[]byte(inputText), 0644)
}
  
func GetBalanceFromTheFile(fileName string) (float64, error){
data, err := os.ReadFile(fileName)

if err != nil {
	return 1000, errors.New("error finding the file")
}
valueText := string(data)
value, err := strconv.ParseFloat(valueText, 64)

if err != nil {
	return 1000, errors.New("error reading from the file")
}
return value, nil
}