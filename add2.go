/*
  asks user to input numbers to add
*/

package main

import "fmt"


func main(){
     fmt.Print("enter first number: \n")
     var num1 float64
     fmt.Scanf("%f",&num1)
     fmt.Print("Enter second number: \n")
      var num2 float64
      fmt.Scanf("%f",&num2)
      fmt.Printf("The sum is: %.2f\n",adder(num1,num2))
}

func adder(a,b float64) float64{
      return a+b
}
