package main
import "fmt"
func main(){
	var num int
	fmt.Print("enter the number of array elements:")
	fmt.Scan(&num)
	even:=make([]int,num)
	fmt.Println("enter", num,"even numbers:")
	for i:=0;i<num;i++{
		fmt.Scan(&even[i])
	}
	fmt.Print("even number are as follows:", even)

	var remove int
	fmt.Print("enter the element to be removed:")
	fmt.Scan(&remove)

	even=append(even[:remove],even[remove+1:]...)
	fmt.Print("after removing the element,slice:",even)
	start,end:=0,0
	fmt.Print("enter the start and the end index which to be removed:")
	fmt.Scan(&start,&end)
	even=append(even[:start],even[end:]...)
	fmt.Println("after removing the multiple elements,slice:",even)
}