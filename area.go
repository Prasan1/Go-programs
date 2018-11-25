/*
* This program calculates the area of 
* a rectangle and the area of a circle from
* the four coordiante points and radius(for circle)
*/
package main

import(
	"fmt"
	"math"
)

//calcualtes the distance form of two coordiante point.
func distance(x1,y1,x2,y2 float64) float64{
      a:= x2-x1
      b:= y2-y1
      return math.Sqrt(a*a + b*b)

}

//function that uses distance to calulate area of a rectangle
func rectangleArea(x1,y1,x2,y2 float64)float64{
     l:= distance(x1,y1,x1,y2)
     w:= distance(x1,y1,x2,y1)
     return l*w

}

// a function that uses math library to calculate area of a circle.
func circleArea(x,y,r float64)float64{
     return math.Pi *r*r

}

func main(){
     var rx1,ry1 float64 = 0,0
     var rx2,ry2 float64 = 10,10
     var cx,cy,cr float64 = 0,0,5


     fmt.Println(rectangleArea(rx1,ry1,rx2,ry2))
     fmt.Println(circleArea(cx,cy,cr))

}
