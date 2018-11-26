/*
 * This program demonstrates the use
 * of struct in go to calulate the
 * area of a rectangle.
*/

package main

import "fmt"

type Rectangle struct{
      w float64
      l float64
}


func rectArea( r Rectangle) float64{
      return  r.w * r.l

}

func main(){
     v:=Rectangle{w:10,l:15}
     fmt.Println(rectArea(v))
}
