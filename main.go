package main

import "fmt"

// julio conchas
// coonchasjimenez@gmail.com

func main()  {
	var n, x, suma int
	var s []int

	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&x)
		s = append(s,x)
	}
	for _, v := range s{
		suma = suma + v
	}
	fmt.Println(suma)
}
