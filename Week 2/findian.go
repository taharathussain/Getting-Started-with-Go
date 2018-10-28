package main
import(
	"fmt"
	"strings"
)
func main(){
	var a string
	fmt.Println("Enter a string with no spaces:")
	fmt.Scan(&a)
	fmt.Println(a)
	i:=strings.ContainsAny(a, "i,a,n")
	if i==true{
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}