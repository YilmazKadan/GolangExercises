package main

import "fmt"

type Kitap struct {
	kitap_id int
	baslik   string
	yazar    string
	konu     string
	sayfa    int
}

func main() {

	// Structer'ları farklı türdeki veri elemanlarını tutmak için kullanırız
	var kitap1 Kitap
	var kitap2 Kitap

	kitap1.kitap_id = 1
	kitap1.baslik = "Kitap 1"
	kitap1.yazar = "Yılmaz Kadan"
	kitap1.konu = "Kitap 1 konusu"
	kitap1.sayfa = 300

	kitap2.kitap_id = 2
	kitap2.baslik = "Kitap 2"
	kitap2.yazar = "Yılmaz Kadan"
	kitap2.konu = "Kitap 2 konusu"
	kitap2.sayfa = 500

	// Kitapların bilgilerini yazdıralım
	bilgiYazdir(kitap1)
	bilgiYazdir(kitap2)

}

// Kitap bilgilerini ekrana yazdıran bir fonksiyon yazalım

func bilgiYazdir(kitap Kitap) {
	fmt.Printf("Kitap ID %d\n", kitap.kitap_id)
	fmt.Printf("Kitap başlığı %s\n", kitap.baslik)
	fmt.Printf("Kitap yazar %s\n", kitap.yazar)
	fmt.Printf("Kitap konu %s\n", kitap.konu)
	fmt.Printf("Kitap sayfa %d\n", kitap.sayfa)
	fmt.Println()
}
