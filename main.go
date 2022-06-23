package main

import (
	"fmt"
	"strings"
)

func manyWords(words ...string) {
	fmt.Println(words)
}

func lenAndUpper(name string) (int, string) {
	defer fmt.Println("end")
	fmt.Println("start")
	return len(name), strings.ToUpper(name)
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		fmt.Println(number)
		total += number
	}
	return total
}

func canIDrinkIf(age int) bool {
	if korAge := age - 2; korAge < 18 {
		return false
	}
	return true
}

func canIDrinkSwitch(age int) bool {
	switch korAge := age - 2; {
	case korAge < 18:
		return false
	case korAge == 18:
		return true
	case korAge > 50:
		return false
	}
	return false
}

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	totLen, _ := lenAndUpper("yumin")
	fmt.Println(totLen)

	manyWords("a", "b", "c")

	fmt.Println(superAdd(1, 2, 3, 4, 5))

	fmt.Println(canIDrinkIf(18))
	fmt.Println(canIDrinkSwitch(20))

	// low level programming
	a := 2
	b := &a
	*b = 10
	fmt.Println(a)

	// array
	names := []string{"a", "b", "c", "d", "e"}
	names[2] = "f"
	names = append(names, "g")
	fmt.Println(names)

	// map
	me := map[string]string{"name": "yumin", "age": "24"}
	for key, value := range me {
		fmt.Println(key, value)
	}

	// struct
	favFood := []string{"coffee", "egg"}
	meStruct := person{name: "yumin", age: 24, favFood: favFood}
	fmt.Println(meStruct)
}
