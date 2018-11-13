/*
  modification to square.go to take a
  number  from a user.
*/

package main
import "fmt"

func main(){
     fmt.Print("Enter a number to square: ")
     var num float64
      fmt.Scanf("%f",&num)
      result:= num * num
      fmt.Printf("The square of %.2f is %.2f\n", num, result)
}
