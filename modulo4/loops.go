package main

import (
	"fmt"
)

func main() {

	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	/*for {
		fmt.Println("Jonas")
		time.Sleep(3 * time.Second)
	}*/

	frutas := []string{"laranja", "maçã", "banana"}
	for i, fruta := range frutas {
		fmt.Println(i)
		fmt.Println(fruta)
	}
}
