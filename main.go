package main
import "fmt"
func num(n **int){
	**n = 8
}


func main(){
	s := 7
	a := &s
	num(&a)
	fmt.Println(s)
	fmt.Print(a)

}