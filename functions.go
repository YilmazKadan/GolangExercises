package main

import "fmt"

func main() {
	// Çıktı: 40
	fmt.Println(max(20, 40))

	a, b := degistir("Kadan", "Yılmaz")

	fmt.Println(a, b)
}

// NOT: Yazılacak olan fonksiyonlar C Dilinde olduğu gibi main fonksiyonu dışına yazılmalıdır.

// Aldığı iki sayıdan büyüğünü geri döndüren bir go fonksiyonu örneği
func max(num1, num2 int) int {
	/* local variable declaration */
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

// ÇOKLU DEĞER DÖNDÜREN GO FONKSİYONU

// Go çoklu değer döndüren fonksiyon özelliğine sahiptir
func degistir(birinci, ikinci string) (string, string) {
	return ikinci, birinci
}
