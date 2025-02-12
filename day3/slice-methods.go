package main
import "fmt"
func main(){
	slice:=[]int{3,6,9}
	fmt.Println("before appending ,slice:",slice)
	//append
	slice=append(slice,12)
	
	slice=append(slice,15,18)
	fmt.Println("after appending new elements,slice:",slice)
	//length of slice
	
	fmt.Println("the length of slice :",len(slice))
	//copy the slice
	new_slice:=make([]int,6)
	copy_slice:=copy(new_slice,slice)
	fmt.Println("the existing slice elements are:",slice)
	fmt.Println("the newly copied slice elements are:",new_slice)
	fmt.Println("the number of elements copied:",copy_slice)
	//clear all elements
	new_slice=nil
	fmt.Println("the new slice is empty:",new_slice)

}