package main

import (
	"fmt"
	"strconv"
)

func main() {

	// KLASİK FOR DÖNGÜSÜ
	for i := 0; i < 20; i++ {
		fmt.Printf("Merhaba " + strconv.Itoa(i) + "\n")
	}

	// WHILE DONGUSU
	n := 1
	for n < 20 {
		n *= 2
	}
	fmt.Println(n)

	// FOREACH DONGUSU
	kelimeler := []string{"merhaba", "dünya"}
	for indis, kelime := range kelimeler {
		fmt.Println(indis, kelime)
	}
	// Strconv.Itoa() metodu bir tip dönüşümü sağlıyor

	// BREAK VE CONTINUE
	// Bildiğiniz gibi bu anahtar kelimeler döngüyü tekrarlamaya veya çıkmaya yarar

	toplam := 0
	for i := 1; i < 5; i++ {
		if i%2 != 0 { // Tek değerleri burada geçiyoruz
			continue
		}
		toplam += i
	}
	fmt.Println(toplam) // 6 (2+4)
	// SONSUZ FOR DONGUSU
	for true {
		fmt.Printf("Bu döngü sonsuza kadar dönecektir \n")
	}
	// Burada true yerine hiçbir değer vermezsek de çalışacaktır

}
