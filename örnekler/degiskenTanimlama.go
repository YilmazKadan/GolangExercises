package main

import "fmt"

func main() {
	var x float64
	x = 20.0
	fmt.Println(x)
	// Tip yazdırmak için aşağıdaki gibi bir kullanım var.
	fmt.Printf("X'in tipi: %T\n", x)

	//  DİNAMİK TÜR BİLDİRİMİ
	//  Burada ':=' ifadesi ile atama yapıldığında, derleyici kendisi tipi anlayacaktır.
	y := 20

	fmt.Printf("Y'nin tipi: %T", y)

	// KARIŞIK TÜR BİLDİRİMİ
	// Aynı anda birden fazla tip bildirimi yapılabilir.Sıra önemli

	var a, b, c = 20, 30, "string"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("a'nın tipi: %T\n", a)
	fmt.Printf("b'nin tipi %T\n", b)
	fmt.Printf("c'nin tipi %T\n", c)

}
