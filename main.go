package main

import (
	"bankamatik/models"
	"fmt"
)

func main() {
	var i int
	var hesapId int
	k := models.Kisi{1, "Cihan", "Bayoğlu", true, nil}

	for i != 0 {
		fmt.Printf("\n*********************************\n")
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
		fmt.Scan(&i)

		if i >= 2 && i <= 6 && len(k.Hesaplar) == 0 {

			fmt.Println("İşlem yapılacak bir hesabınız bulunmamaktadır.")

		} else if i >= 7 && i <= 10 && len(k.Hesaplar) < 2 {

			fmt.Println("Bu işlemi gerçekleştirmek için en az 2 hesaba ihtiyacınız vardır.")

		} else {

			switch i {
			case 1:
				var hesapTuru int
				fmt.Println("Oluşturmak istediğiniz hesap türünü seciniz")
				fmt.Println("1-)tl -- 2-)dolar -- 3-)euro")
				fmt.Scan(&hesapTuru)
				if hesapTuru < 0 || hesapTuru > 3 {
					fmt.Printf("Hatalı hesap türü seçtiniz.")
				} else {
					hesapId++
					k.HesapOlustur(hesapId, hesapTuru)
				}

				break
			case 2:

				var e int // var olan hesapların id'lerini görmek için kullanılacak.
				for e = 0; e < len(k.Hesaplar); e++ {
					fmt.Printf("%d. Hesap , ID: %d , Durum: %t , Tür: %d , Bakiye: %.2f\n",
						e, k.Hesaplar[e].ID, k.Hesaplar[e].Durum, k.Hesaplar[e].Tur, k.Hesaplar[e].Bakiye)
				}

				println("Blokaj işlemi yapılacak hesabı giriniz.")
				fmt.Scan(&e)
				var opsiyon bool

				println("Bloke etmek için 0, blokeyi kaldırmak için 1 giriniz.")
				fmt.Scan(&opsiyon)
				k.HesapBlokaj(&k.Hesaplar[e], opsiyon)
				break
			case 3:

				var e int
				var tutar float64 // yatırılacak tutarı almak için kullanılacak.
				for e = 0; e < len(k.Hesaplar); e++ {
					fmt.Printf("%d. Hesap , ID: %d , Durum: %t , Tür: %d , Bakiye: %.2f\n",
						e, k.Hesaplar[e].ID, k.Hesaplar[e].Durum, k.Hesaplar[e].Tur, k.Hesaplar[e].Bakiye)
				}
				println("Para yatırılacak hesabı giriniz")
				fmt.Scan(&e)
				println("Yatırılacak tutarı giriniz.")
				fmt.Scan(&tutar)
				k.ParaYatir(&k.Hesaplar[e], tutar)

				break
			case 4:
				//para çek

				var e int
				var tutar float64 // çekilecek tutarı almak için kullanılacak.
				for e = 0; e < len(k.Hesaplar); e++ {
					fmt.Printf("%d. Hesap , ID: %d , Durum: %t , Tür: %d , Bakiye: %.2f\n",
						e, k.Hesaplar[e].ID, k.Hesaplar[e].Durum, k.Hesaplar[e].Tur, k.Hesaplar[e].Bakiye)
				}
				println("Para çekilecek hesabı giriniz")
				fmt.Scan(&e)
				println("Çekilecek tutarı giriniz.")
				fmt.Scan(&tutar)
				k.ParaCek(&k.Hesaplar[e], tutar)
				break
			case 5:
				var e int // var olan hesapların id'lerini görmek için kullanılacak.
				for e = 0; e < len(k.Hesaplar); e++ {
					fmt.Printf("%d. Hesap , ID: %d , Durum: %t , Tür: %d , Bakiye: %.2f\n",
						e, k.Hesaplar[e].ID, k.Hesaplar[e].Durum, k.Hesaplar[e].Tur, k.Hesaplar[e].Bakiye)
				}
				break
			case 6:
				//hesapsil

				var e int
				for e = 0; e < len(k.Hesaplar); e++ {
					fmt.Printf("%d. Hesap , ID: %d , Durum: %t , Tür: %d , Bakiye: %.2f\n",
						e, k.Hesaplar[e].ID, k.Hesaplar[e].Durum, k.Hesaplar[e].Tur, k.Hesaplar[e].Bakiye)
				}
				println("Silmek istediğiniz hesabı giriniz")
				fmt.Scan(&e)
				if !k.Hesaplar[e].Durum {
					println("Bloke hesap silinemez.")

				} else {
					k.Hesaplar = append(k.Hesaplar[:e], k.Hesaplar[e+1:]...)
					fmt.Printf("%d. hesabınız silindi.", e)
				}

				break
			case 7:
				//Döviz al (tl hesabından usd veya eur hesabına {para gönder : case3} )

				var e, e2 int     // var olan hesapların id'lerini görmek için kullanılacak.
				var tutar float64 // gönderilecek tutarı almak için kullanılacak.
				for e = 0; e < len(k.Hesaplar); e++ {
					fmt.Printf("%d. Hesap , ID: %d , Durum: %t , Tür: %d , Bakiye: %.2f\n",
						e, k.Hesaplar[e].ID, k.Hesaplar[e].Durum, k.Hesaplar[e].Tur, k.Hesaplar[e].Bakiye)
				}
				println("döviz hesabı giriniz")
				fmt.Scan(&e)
				println("tl hesabı giriniz")
				fmt.Scan(&e2)
				println("Döviz alınacak tutarı giriniz (tl)")
				fmt.Scan(&tutar)
				k.DovizAl(&k.Hesaplar[e], &k.Hesaplar[e2], tutar)

				break
			case 8:
				//Döviz bozdur (usd veya eur hesabından tl hesabona {para gönder: case3} )

				var e, e2 int     // var olan hesapların id'lerini görmek için kullanılacak.
				var tutar float64 // gönderilecek tutarı almak için kullanılacak.
				for e = 0; e < len(k.Hesaplar); e++ {
					fmt.Printf("%d. Hesap , ID: %d , Durum: %t , Tür: %d , Bakiye: %.2f\n",
						e, k.Hesaplar[e].ID, k.Hesaplar[e].Durum, k.Hesaplar[e].Tur, k.Hesaplar[e].Bakiye)
				}
				println("döviz hesabı giriniz")
				fmt.Scan(&e)
				println("tl hesabı giriniz")
				fmt.Scan(&e2)
				println("Bozdurulacak döviz tutarı tutarını giriniz ")
				fmt.Scan(&tutar)
				k.DovizBoz(&k.Hesaplar[e], &k.Hesaplar[e2], tutar)
				break
			case 9:

				var e, e2 int     // var olan hesapların id'lerini görmek için kullanılacak.
				var tutar float64 // gönderilecek tutarı almak için kullanılacak.
				for e = 0; e < len(k.Hesaplar); e++ {
					fmt.Printf("%d. Hesap , ID: %d , Durum: %t , Tür: %d , Bakiye: %.2f\n",
						e, k.Hesaplar[e].ID, k.Hesaplar[e].Durum, k.Hesaplar[e].Tur, k.Hesaplar[e].Bakiye)
				}
				println("Alıcı hesabı giriniz")
				fmt.Scan(&e)
				println("Gönderici hesabı giriniz")
				fmt.Scan(&e2)
				println("Gönderilecek tutarı giriniz")
				fmt.Scan(&tutar)

				k.ParaGonder(&k.Hesaplar[e], &k.Hesaplar[e2], tutar)
				break
			case 10:
				var e, e2 int // var olan hesapların id'lerini görmek için kullanılacak.
				for e = 0; e < len(k.Hesaplar); e++ {
					fmt.Printf("%d. Hesap , ID: %d , Durum: %t , Tür: %d , Bakiye: %.2f\n",
						e, k.Hesaplar[e].ID, k.Hesaplar[e].Durum, k.Hesaplar[e].Tur, k.Hesaplar[e].Bakiye)
				}
				println("ana hesabı giriniz")
				fmt.Scan(&e)
				println("İçeriği ana hesaba aktarılacak hesabı giriniz. (Bu hesap işlem sonunda silinecektir.)")
				fmt.Scan(&e2)
				k.HesapBirlestir(&k.Hesaplar[e], &k.Hesaplar[e2], e2)
				break
			default:
				fmt.Printf("Lütfen 1-10 arası seçim yapınız.")
				break
			}
		}
	}

}
