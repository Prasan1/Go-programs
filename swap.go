/*
  this program swaps two int. Also, go func
   can return two values and this program 
    demonstrates that. 
*/


package main

func main(){
    a,b:=swap(2,5)
    println(a,b)
}

func swap(a,b int) (int,int){
      return b,a
}
