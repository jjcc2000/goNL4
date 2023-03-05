package main

import("fmt")

func main(){
	ej1()

}

func ej1(){
	x :=[5]int{1,2,3,4,5}

	for i, v:= range x{
		fmt.Printf("The position is: %v the value is: %v\n",  i,v)
	}
	fmt.Printf("The type of the array is: %T\n",x)
}