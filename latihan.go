package main
import "fmt"

func kaprekar(n int)int{
	var terbesar, terkecil, i int
	terbesar = n%10
	terkecil = n%10
	for n > 0{
		i = n%10
		if i > terbesar{
			terbesar = i
		}
		if i < terkecil{
			terkecil = i
		}
		n = n/10
	}
	return terbesar-terkecil
}
func main(){
	var num int
	fmt.Scan(&num)
	fmt.Print(kaprekar(num), " Langkah")
}

