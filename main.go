package main

import (
	"bankamatik/models"
	"fmt"
)

func main() {
	i := 1            // Yapılacak işlem numarasını tutar.
	var e int         // Fonksiyonun 1. değişken değerini tutar.
	var e2 int        // Fonksiyonun 2. değişken değerini tutar.
	var tutar float64 // Tutar parametresi alan fonksiyona gönderilecek parametreyi tutar.
	k := models.Kisi{1, "Cihan", "Bayoğlu", true, nil}

	for i != 0 {
		println("\n*********************************\n")
		println("Yapmak istediğiniz işlemi giriniz")
		println("1-)  Hesap oluştur")         // 0 hesap gerekli
		println("2-)  Hesap bloke işlemleri") // 1hesap gerekli
		println("3-)  Para yatır")            // 1 hesap gerekli
		println("4-)  Para çek")              // 1 hesap gerekli
		println("5-)  Bakiye göster")         // 1 hesap gerekli
		println("6-)  Hesap sil")             // 1 hesap gerekli
		println("7-)  Döviz al")              // 2 hesap gerekli
		println("8-)  Döviz bozdur")          // 2 hesap gerekli
		println("9-)  Para gönder")           // 2 hesap gerekli
		println("10-) Hesap birleştirme")     // 2 hesap gerekli
		println("11-)Çıkış  için 0 giriniz.")
		println("\n*********************************\n")
		fmt.Scan(&i)

		if i >= 2 && i <= 6 && len(k.Hesaplar) == 0 {
			fmt.Println("İşlem yapılacak bir hesabınız bulunmamaktadır.")

		} else if i >= 7 && i <= 10 && len(k.Hesaplar) < 2 {
			fmt.Println("Bu işlemi gerçekleştirmek için en az 2 hesaba ihtiyacınız vardır.")

		} else {
			switch i {
			case 0:
				fmt.Printf("Çıkış yaptınız.")
				break

			case 1:
				k.HesapOlustur(&k)
				break

			case 2:
				k.BakiyeGoster()

				k.HesapBlokaj()

				break

			case 3:
				k.BakiyeGoster()

				println("Para yatırılacak hesabı giriniz")
				fmt.Scan(&e)
				println("Yatırılacak tutarı giriniz.")
				fmt.Scan(&tutar)

				if models.HataliMi(&k, e) {
					println("Yanlış hesap girdiniz.")
				} else {
					k.ParaYatir(&k.Hesaplar[e], tutar)
				}
				break

			case 4: //para çek
				k.BakiyeGoster()

				println("Para çekilecek hesabı giriniz")
				fmt.Scan(&e)
				println("Çekilecek tutarı giriniz.")
				fmt.Scan(&tutar)

				if models.HataliMi(&k, e) {
					println("Yanlış hesap girdiniz.")
				} else {
					k.ParaCek(&k.Hesaplar[e], tutar)
				}
				break

			case 5:
				k.BakiyeGoster()
				break

			case 6: //hesapsil
				k.BakiyeGoster()

				println("Silmek istediğiniz hesabı giriniz")
				fmt.Scan(&e)
				if models.HataliMi(&k, e) {
					println("Yanlış hesap girdiniz.")

				} else if !k.Hesaplar[e].Durum {
					println("Bloke hesap silinemez.")

				} else {
					k.Hesaplar = append(k.Hesaplar[:e], k.Hesaplar[e+1:]...)
					fmt.Printf("%d. hesabınız silindi.", e)
				}

				break
			case 7: //Döviz al (tl hesabından usd veya eur hesabına {para gönder : case3} )
				k.BakiyeGoster()

				println("döviz hesabı giriniz")
				fmt.Scan(&e)
				println("tl hesabı giriniz")
				fmt.Scan(&e2)
				println("Döviz alınacak tutarı giriniz (tl)")
				fmt.Scan(&tutar)
				if models.HataliMi(&k, e, e2) {
					println("Yanlış hesap girdiniz.")
				} else {
					k.DovizAl(&k.Hesaplar[e], &k.Hesaplar[e2], tutar)
				}
				break

			case 8: //Döviz bozdur (usd veya eur hesabından tl hesabona {para gönder: case3} )
				k.BakiyeGoster()

				println("döviz hesabı giriniz")
				fmt.Scan(&e)
				println("tl hesabı giriniz")
				fmt.Scan(&e2)
				println("Bozdurulacak döviz tutarı tutarını giriniz ")
				fmt.Scan(&tutar)

				if models.HataliMi(&k, e, e2) {
					println("Yanlış hesap girdiniz.")
				} else {
					k.DovizBoz(&k.Hesaplar[e], &k.Hesaplar[e2], tutar)
				}
				break

			case 9:
				k.BakiyeGoster()

				println("Alıcı hesabı giriniz")
				fmt.Scan(&e)
				println("Gönderici hesabı giriniz")
				fmt.Scan(&e2)
				println("Gönderilecek tutarı giriniz")
				fmt.Scan(&tutar)

				if models.HataliMi(&k, e, e2) {
					println("Yanlış hesap girdiniz.")
				} else {
					k.ParaGonder(&k.Hesaplar[e], &k.Hesaplar[e2], tutar)
				}

				break

			case 10:
				k.BakiyeGoster()

				println("ana hesabı giriniz")
				fmt.Scan(&e)
				println("İçeriği ana hesaba aktarılacak hesabı giriniz. (Bu hesap işlem sonunda silinecektir.)")
				fmt.Scan(&e2)

				if models.HataliMi(&k, e, e2) {
					println("Yanlış hesap girdiniz.")
				} else {
					k.HesapBirlestir(&k.Hesaplar[e], &k.Hesaplar[e2], e2)
				}

				break

			default:
				fmt.Println("Lütfen 1-10 arası seçim yapınız.")
				break
			}
		}
	}

}
