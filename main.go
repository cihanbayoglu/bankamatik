package main

import (
	"bankamatik/models"
	"fmt"
	"log"
	"os"
)

var Operations = [11]string{
	"0-)  Çıkış",
	"1-)  Hesap oluştur",
	"2-)  Hesap bloke işlemleri",
	"3-)  Para yatır",
	"4-)  Para çek",
	"5-)  Bakiye göster",
	"6-)  Hesap sil",
	"7-)  Döviz al",
	"8-)  Döviz bozdur",
	"9-)  Para gönder",
	"10-) Hesap birleştirme",
}

func main() {
	i := 1              // Yapılacak işlem numarasını tutar.
	var var1 int        // Fonksiyonun 1. değişken değerini tutar.
	var var2 int        // Fonksiyonun 2. değişken değerini tutar.
	var value float64   // value parametresi alan fonksiyona gönderilecek parametreyi tutar.
	isHappened := false // log için işlemin gerçekleşip gerçekleşmediğini kontrol eder.
	k := models.Kisi{1, "Cihan", "Bayoğlu", true, nil}

	// log
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)

	for i != 0 {
		println("\n*********************************\n")
		for operation := range Operations {
			println(Operations[operation])
		}
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
				isHappened = true
				break

			case 1:
				isHappened = k.HesapOlustur(&k)
				break

			case 2:
				k.BakiyeGoster()

				isHappened = k.HesapBlokaj()

				break

			case 3:
				k.BakiyeGoster()

				println("Para yatırılacak hesabı giriniz")
				fmt.Scan(&var1)
				println("Yatırılacak tutarı giriniz.")
				fmt.Scan(&value)

				if models.HataliMi(&k, var1) {
					println("Yanlış hesap girdiniz.")
				} else {
					isHappened = k.ParaYatir(&k.Hesaplar[var1], value)
				}
				break

			case 4: //para çek
				k.BakiyeGoster()

				println("Para çekilecek hesabı giriniz")
				fmt.Scan(&var1)
				println("Çekilecek tutarı giriniz.")
				fmt.Scan(&value)

				if models.HataliMi(&k, var1) {
					println("Yanlış hesap girdiniz.")
				} else {
					isHappened = k.ParaCek(&k.Hesaplar[var1], value)
				}
				break

			case 5:
				isHappened = k.BakiyeGoster()
				break

			case 6: //hesapsil
				k.BakiyeGoster()

				println("Silmek istediğiniz hesabı giriniz")
				fmt.Scan(&var1)
				if models.HataliMi(&k, var1) {
					println("Yanlış hesap girdiniz.")

				} else if !k.Hesaplar[var1].Durum {
					println("Bloke hesap silinemez.")

				} else {
					k.Hesaplar = append(k.Hesaplar[:var1], k.Hesaplar[var1+1:]...)
					fmt.Printf("%d. hesabınız silindi.", var1)
					isHappened = true
				}

				break
			case 7: //Döviz al (tl hesabından usd veya eur hesabına {para gönder : case3} )
				k.BakiyeGoster()

				println("döviz hesabı giriniz")
				fmt.Scan(&var1)
				println("tl hesabı giriniz")
				fmt.Scan(&var2)
				println("Döviz alınacak tutarı giriniz (tl)")
				fmt.Scan(&value)
				if models.HataliMi(&k, var1, var2) {
					println("Yanlış hesap girdiniz.")
				} else {
					isHappened = k.DovizAl(&k.Hesaplar[var1], &k.Hesaplar[var2], value)
				}
				break

			case 8: //Döviz bozdur (usd veya eur hesabından tl hesabona {para gönder: case3} )
				k.BakiyeGoster()

				println("döviz hesabı giriniz")
				fmt.Scan(&var1)
				println("tl hesabı giriniz")
				fmt.Scan(&var2)
				println("Bozdurulacak döviz tutarı tutarını giriniz ")
				fmt.Scan(&value)

				if models.HataliMi(&k, var1, var2) {
					println("Yanlış hesap girdiniz.")
				} else {
					isHappened = k.DovizBoz(&k.Hesaplar[var1], &k.Hesaplar[var2], value)
				}
				break

			case 9:
				k.BakiyeGoster()

				println("Alıcı hesabı giriniz")
				fmt.Scan(&var1)
				println("Gönderici hesabı giriniz")
				fmt.Scan(&var2)
				println("Gönderilecek tutarı giriniz")
				fmt.Scan(&value)

				if models.HataliMi(&k, var1, var2) {
					println("Yanlış hesap girdiniz.")
				} else {
					isHappened = k.ParaGonder(&k.Hesaplar[var1], &k.Hesaplar[var2], value)
				}

				break

			case 10:
				k.BakiyeGoster()

				println("ana hesabı giriniz")
				fmt.Scan(&var1)
				println("İçeriği ana hesaba aktarılacak hesabı giriniz. (Bu hesap işlem sonunda silinecektir.)")
				fmt.Scan(&var2)

				if models.HataliMi(&k, var1, var2) {
					println("Yanlış hesap girdiniz.")
				} else {
					isHappened = k.HesapBirlestir(&k.Hesaplar[var1], &k.Hesaplar[var2], var2)
				}

				break

			default:
				fmt.Println("Lütfen 1-10 arası seçim yapınız.")
				break
			}
		}
		if isHappened {
			log.Println(Operations[i])
		}

	}

}
