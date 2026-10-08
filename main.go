package main

import (
	"fmt"
    "github.com/01-edu/z01"
	)


func hello(s string){
  for _, n := range s {
	z01.PrintRune(n)  // this is a new way for me inestand of doing this look down 
	z01.PrintRune('\n')


  }
}

func main(){
	s := "hello world"
	hello(s)
	
}

// $$$$$$$$$$$$$$$$$$$       // $$$$$$$$$$$$$$$$$$$$$         //$$$$$$$$$$$$$$$$$$$$$$$$$$$$
func main (){
	k := "hello world"
	for _, n := range k {
		fmt.Print(string(n))  // this is a strateforword way
	}
}