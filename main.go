package main
import (
	"fmt"
	"strings"
)


// those are shortcuts

func main(){
	myname := "hamo salem"
	fmt.Println(strings.Contains(myname,"hamo")) // its clear from the word it makes sure is this thing exict or not and it gives you result eather true or fulse at the end
	fmt.Println(strings.Contains(myname,"omar"))
	fmt.Println(strings.ReplaceAll(myname , "salem", "abdelghany"))  // the same though but this use it to change a word or something and don't forget its a temporary thing
	fmt.Println(strings.ToUpper(myname)) // just makes everything C like this  (HAMO SALEM)
	fmt.Print("the main one ",myname)   // so here is the bottom line these shortcuts do not change forever just for a certin time in the moment
}



//$$$$$$$$$$$$$$$$$$$	//$$$$$$$$$$$$$$$$$$$	//$$$$$$$$$$$$$$$$$$$	//$$$$$$$$$$$$$$$$$$$ //$$$$$$$$$$$$$$$$$$$


func names(b int , u int)(int,int){ // i didn't give them names here as you see but its ok i called them n and j down here 
	return b + u , b - u
}

func main(){
	j := 10
	k := 5
	n,g := names(j,k) // !!
	fmt.Printf("this is total %d  and this is the minse %d ",n , g) // new thing here using %d inestand of %v 
}

//$$$$$$$$$$$$$$$$$$$	//$$$$$$$$$$$$$$$$$$$	//$$$$$$$$$$$$$$$$$$$	//$$$$$$$$$$$$$$$$$$$ //$$$$$$$$$$$$$$$$$$$
func str(n string) (i string ){
	return n
}

func main(){

	i := str("hello world") // str("hello world") this is concidered n ok 
	fmt.Print(i) // this is (i string ){ return n }  i have to print i to get what n has

}
//$$$$$$$$$$$$$$$$$$$	//$$$$$$$$$$$$$$$$$$$	//$$$$$$$$$$$$$$$$$$$	//$$$$$$$$$$$$$$$$$$$ //$$$$$$$$$$$$$$$$$$$



func names(p,b int )(total int,mines int) {   // i can also write it like this names(p,b int ) not names(p int, b int )
	// return p + b , p - b // two apporation so two (total int,mines int) got it ?
}

func main(){
	total,mines:= names(10,5) // directly inestand of creating two varaibles  names(10,5)
	fmt.Printf("this is total %d and this is the minse %d ",total,mines) // fmt.Print( %d ) > new thing i think it's used for numbers 
}




//$$$$$$$$$$$$$$$$$$$	//$$$$$$$$$$$$$$$$$$$	//$$$$$$$$$$$$$$$$$$$	//$$$$$$$$$$$$$$$$$$$ //$$$$$$$$$$$$$$$$$$$

func listofname(n string,l int) (i string , o int){   // i want you to picture this kinda apporation i mean the soucend one 
	return n , l  // but the first is porameters you either give them number or name ok you can do an apporation with it but i think just with pointers
}
func main(){
	i,o := listofname("abdo mostafa", 18)  // i can give it a name from the first moment or from here its ok = (i,o :=) 
	fmt.Printf("the name of the man is %v and his age is %d \n",i,o) //  listofname("abdo mostafa", 18) i can create a varaible and give a number or directory like this 
	fmt.Println(strings.ToUpper(i)) 
}