package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	fmt.Println("Random -->> ", rand.Float32())
	fmt.Println("Random -->> ", rand.Float64())

	// literally
	fmt.Println("Hello World -->> ")
	fmt.Println(20 + 20)

	//type const
	const price float32 = 275.00
	const tax float32 = 27.50
	fmt.Println("summ -->>", price+tax)

	// not type const
	const price2, tax2 = 275.00, 666
	const quantity, inStock = 2, true // не указываем тип
	fmt.Println("summ -->>", quantity*(price2+tax2), inStock)

	// pointers2
	names := [3]string{"alice", "charlie", "bob"}

	secondName := &names[1]
	fmt.Println(*secondName)
	sort.Strings(names[:])
	fmt.Print(*secondName)

}
