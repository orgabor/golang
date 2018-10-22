package main


// YOUR CODE HERE

import(
"fmt"

)

func main(){

fmt.Println(CalculateTax(280.0, 0.27))
} 

func CalculateTax (x,y float64) float64{
  return x * y 
}