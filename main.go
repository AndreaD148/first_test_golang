package main

import "fmt"	

func changeVal(a *int) {
	*a = 100;
}

func swap(a *int, b *int) {
	flag := *a
	*a = *b
	*b = flag

}

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


	//////////////////////////////////////////////////////////
	// using function that take a pointer of two			//
	// variables and chnage the value of a passing variable	//
	// and after, in the main function, printing the result	//
	//////////////////////////////////////////////////////////

	num_to_change := 2

	fmt.Println("Now the value of num_to_change is: ", num_to_change)

	fmt.Println("Now is gonna be changing the value?")
	
	changeVal(&num_to_change)

	fmt.Println(num_to_change)

	//////////////////////////////////////////////////////////
	// Now I'm trying to swap to values						//
	//////////////////////////////////////////////////////////

	value1 := 1
	value2 := 6

	fmt.Println("The value of value1 and value 2 are: ", value1, value2)

	swap(&value1, &value2)

	fmt.Println("Now that I swapped those two varables: ", value1, value2)



}