/*
  argument reader
  this simply prints what is in Args[0]
*/

package main
import "fmt"
import "os"

func main(){
     fmt.Println("You entered: ",os.Args[0])
}
