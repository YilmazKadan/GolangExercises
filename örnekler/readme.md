# Burada öğrenirken aldığım notlar bulunmaktadır.

## Projeyi ayağa kaldırmak

Bir go dosyasını ayağa kaldırmak için çalıştırılmak istenilen dosyanın bulunduğu dizine girilip <pre> go run dosya_adi.go </pre> şeklinde bir kullanım var.


## Sabitler

Sabitler program yürütülürken değiştirilemeyen sabit değerleri ifade ederler.
Bu değişmeyen değerlere sabitler(constans) denir.
İlgili örneklerin dosyası : [bu linke tıklayarak gidebilirsiniz](sabitler.go)

## Escape(Kaçış) karakterleri

Diğer programlama dillerinde olduğu gibi burada da kaçış karakterleri ve kullanım amaçları vardır. En çok kullandığımız kaçış karakteri amaclarından ikisi ""(çift tırnak) veya alt satıra inme için "\n" komutlarıdır. Kaçış karakteri "\"(Ters slaş) ile başlar. Derleyici bunu gördüğü yerde ardından gelen karakter ile ne yapmaya çalışıldığını anlayacaktır.

|     **\\**     |                      \ Karakteri                     |
|:--------------:|:----------------------------------------------------:|
|     **\'**     |                      ' Karakteri                     |
|     **\"**     |                      " karakteri                     |
|     **\?**     |                      ? karakteri                     |
|     **\a**     |                    Uyarı veya zil                    |
|     **\b**     |                 Backspace(geri alma)                 |
|     **\f**     |                Form feed(form besleme)               |
|     **\n**     |                  Newline(Yeni Satır)                 |
|     **\r**     |                    Carriage return                   |
|     **\t**     |                       Yatay tab                      |
|     **\v**     |                       Dikey Tab                      |
|    **\ooo**    |           Bir ila üç basamaklı sekizli sayı          |
| **\xhh . . .** | Bir veya daha fazla basamaktan oluşan onaltılık sayı |

## Döngüler

Go dilinde döngülerin hepsi **for** anahtar sözcüğü ile başlar burada for döngüsü while, do while, foreach ve klasik for döngüsü şeklinde hareket edebiliyor. Bu sayede çok amaçlı bir for tipi elde etmişler.