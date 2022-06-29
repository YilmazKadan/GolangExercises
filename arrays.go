package main

import "fmt"

func main() {
	var n [10]int /* 10 Elemanlı bir tam sayısı dizisi tanımlar */

	// Tüm elemanlara sırasıyla 1'den 10'a kadar sayıların 10 fazlasını ekleyelim

	for i := 0; i < 10; i++ {
		n[i] = i + 10
	}

	// Elemanları ekrana yazdıralım

	for j := 0; j < 10; j++ {
		fmt.Println(n[j])
	}
}
