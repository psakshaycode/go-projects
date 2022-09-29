package main
import "fmt"
//this is a comment
func main(){
x := [5]float64{23,2,1,5,6}

var total float64 =0
for _,value := range x{
	//fmt.Println(i)
	total +=value
}
fmt.Println(total)
}