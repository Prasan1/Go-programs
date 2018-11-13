/*
  This program reads in the argument
  entered by the user
*/

package main

import(
    "fmt"
     "os"
)

 func main(){
    fmt.Printf("You entered %s\n",os.Args[1:])

}
