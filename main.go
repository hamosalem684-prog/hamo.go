package main 
import "fmt"

func names(p,b int )(total int,mines int) {   // i can also write it like this names(p,b int ) not names(p int, b int )
	return p + b , p - b // two apporation so two (total int,mines int) got it ?
}

func main(){
	fmt.Print(names(10,5)) 
}