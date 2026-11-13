

package main
import (
	"fmt"
	"errors"
)

type student struct {
	fullname string
	city string
	qualty string
	username string
	age,id int 
	grades []int
}



func main(){
	result := student{
	fullname: "mohamed salem",
	username: "hamo123",
	city: "USA",
	qualty: "great and talanted",
	age: 18,
	id: 110,
	grades: []int{70,80,90,90},
    }
	theaddress := "The Original Things"
	fmt.Println(theaddress)
    fmt.Println(result.fullname)
	fmt.Println(result.username)
	fmt.Println(result.city)
	fmt.Println(result.qualty)
	fmt.Println(result.age)
	fmt.Println(result.id)
	fmt.Println(result.grades) // [70,80,90,90]
    
    label := "after modification"
    fmt.Println(label)


	fmt.Printf("his name is [%v]",result.fullname)
	fmt.Printf("the username is [%v]\n",result.username)
	username1(result.username)
	fmt.Print("\n")
	fmt.Printf("his from [%v]\n",result.city)
	fmt.Printf("the behaver is [%v]\n",result.qualty)
	massage , err := modefytheage(result.age)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(massage)
	fmt.Printf("the id is[%d]\n",result.id)
	checking(result.grades)
	fmt.Print("\n")
}

func upper(s string)string{
	// mohamed salem
	var i string
	for a := 0 ; a < len(s) ; a++{
		if (a == 0 || s[a-1] == ' ')&&
		   (s[a] >= 'a' && s[a] <= 'z'){
			i += string(s[a]- 32) 
		}else{
			i += string(s[a])
		} 
	}
	return i
}

func username1(s string) {
	//"hamo123"
	valid := len(s) > 0
	for _, a := range s {
		if !(a >= 'a' && a <= 'z')&&
		   !(a >= '0' && a <= '9'){
			valid = false 
			break
		}
	}
	if valid {
		fmt.Println("valid username")
	}else{
		fmt.Println("invalid username")
	}
}

func modefytheage(s int)(string,error){
	if s < 18 {
		return "",errors.New("we accept at least 18")
	}
	return "accepted",nil
}

func checking(s []int){ 
	for _, a := range s {
		if a <= 50 || a <= 40 { // {70,80,90,90}
			fmt.Println("droped one in grades [%d]",a)
		}
		if a > 100 {
			break
		}
	}
}

