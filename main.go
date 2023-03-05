package main

import (
	"fmt"
	
)

func main(){
	ej10()
}

func ej1(){
	x :=[5]int{1,2,3,4,5}

	for i, v:= range x{
		fmt.Printf("The position is: %v the value is: %v\n",  i,v)
	}
	fmt.Printf("The type of the array is: %T\n",x)
}
func ej2(){
	x := []int {0,1,2,3,4,5,6,7,8,9}
	for  v :=range x{
		fmt.Printf("The value is: %v\n",v)
	}
	fmt.Printf("The type of the Slice is: %T\n",x)
}
func ej3(){
	x := []int {0,1,2,3,4,5,6,7,8,9}
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
}
func ej4(){
	x:= []int{42,43,44,45,46,47,48,49,50,51}
 	x = append(x, 52)
	fmt.Println(x)
	x= append(x, 53,54,55)
	fmt.Println(x)
	y:= []int{56,57,58,59,60}
	x = append(x, y...)
	fmt.Println(x)
}
func ej5(){
	x:= []int{42,43,44,45,46,47,48,49,50,51}
	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}
func ej6(){
	x:= make([]string,50, 50)
	fmt.Println(cap(x))
	fmt.Println(len(x))
	
	states:= []string{"Alabama" ,"Alaska", "Arizona", "Arkansas", "California", "Colorado"," Connecticut", "Delaware", "Florida"," Georgia"," Hawaii","Idaho", "Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota", "Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico", "New York", "North Carolina", "North Dakota", "Ohio", "Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina", "South Dakota", "Tennessee", "Texas", "Utah", "Vermont", "Virginia", "Washington", "West Virginia", "Wisconsin",  "Wyoming."}

	for i,v :=range states {
		x[i]=v
	}
	fmt.Println("This is after the assigments")
	fmt.Println(cap(x))
	fmt.Println(len(x))
	fmt.Println(x)
}

func ej7(){
	y := []string {"James","Bond","Shaken","Not Stired"}
	z:= []string {"Miss", "MoneyPenny", "Hello","James"}
	x := [][]string{y,z}
	
	for i, v:= range x{
		fmt.Printf("This is the column: %v\n", i)
		for p, val :=range v{
			fmt.Printf("\tThe position in the row: %v The row value is: %v\n",p,val)
		} 
	}
}

func ej8(){
	m := map[string][]string{
		"Johan": []string {"WorkOut", "Food","Music"},
		"Hannia":[]string{"Sleep","Driving","Food"} ,
		"Snoop": []string{"Run","Drink","Bark"},
	}
	for k, val := range m{
		fmt.Printf("The key value is:%v\n",k)
		for i,s := range val{
			fmt.Printf("\tThe position is :%v The value of the position is:%v\n",i,s)
		}
	}
}

func ej9(){
	m := map[string][]string{
		"Johan": []string {"WorkOut", "Food","Music"},
		"Hannia":[]string{"Sleep","Driving","Food"} ,
		"Snoop": []string{"Run","Drink","Bark"},
	}

	m["asd"]= []string{"Guitar","Piano","bass"}

	for k, val := range m{
		fmt.Printf("The key value is:%v\n",k)
		for i,s := range val{
			fmt.Printf("\tThe position is :%v The value of the position is:%v\n",i,s)
		}
	}
}
func ej10(){
	m := map[string][]string{
		"Johan": []string {"WorkOut", "Food","Music"},
		"Hannia":[]string{"Sleep","Driving","Food"} ,
		"Snoop": []string{"Run","Drink","Bark"},
	}
	delete(m,"Johan")

	for k, val := range m{
		fmt.Printf("The key value is:%v\n",k)
		for i,s := range val{
			fmt.Printf("\tThe position is :%v The value of the position is:%v\n",i,s)
		}
	}
}