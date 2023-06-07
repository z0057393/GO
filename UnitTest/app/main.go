package main 

import (
	"fmt"
)

func main (){
	var index int = 1
	for index <= 100{
		fmt.Println(check(index))
		index++
	}
}

func check(num int) string{
	if num % 3 == 0 && num % 5 == 0 {
		return FizzBuzz()
	} else if num % 3 == 0 {
		return Fizz()
	} else if num % 5 == 0 {
		return Buzz()
	}else{
		return "Fizz"
	}
}

func Fizz() string{

	return "Fizz"
}

func Buzz() string {
	return "Buzz"
}

func FizzBuzz() string{
	return "FizzBuzz"
}