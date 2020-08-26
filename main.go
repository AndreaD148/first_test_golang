package main

import "fmt"	



func main() {
	var x = 5
	y := 2

	var z *int = &x


	fmt.Println(x, y)
	fmt.Println(x, *z)
	

	//change the value of x using his pointer
	*z = 4

	fmt.Println(x, *z)

	var my_array []int = []int{1, 2, 3, 4}
	
	fmt.Println(x, my_array)
	
	copy_of_array := my_array
	copy_of_array[0] = 100
	
	fmt.Println(my_array, copy_of_array)
}