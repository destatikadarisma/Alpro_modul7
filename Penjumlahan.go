package main
import "fmt"
func main(){
	var x, y, j, jumlah  int
	fmt.Scan(&x,&y)
	for j = x; j <= y; j++{
		jumlah = jumlah + j
	}
	fmt.Println("hasil penjulmahan", jumlah)
}