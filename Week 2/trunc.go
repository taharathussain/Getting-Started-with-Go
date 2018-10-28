package main
import(
	"fmt"
)
func main(){
	var x float64
	fmt.Println("Enter First float number:")
	fmt.Scan(&x)
		y := fmt.Sprintf("%.0f", x)
	
	var a float64
	fmt.Println("Enter Second float number:")
	fmt.Scan(&a)
		b := fmt.Sprintf("%.0f", a)
		fmt.Println("First truncated number is:",y)
		fmt.Println("Second truncated number is:",b)
}