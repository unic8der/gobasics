package main

import (
	"fmt"
	"time"
)

func main() {
	myType := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Bool my mind!!!")
		case int:
			fmt.Println("Int to you!!!")
		default:
			fmt.Printf("Definitely not my type %T\n", t)
		}
	}

	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Breakfast time!!!")
		myType(hour)
	case hour < 17:
		fmt.Println("Snack time!!!")
		myType(hour)
	default:
		fmt.Println("Party time!!!")
		myType(hour)
	}

}
