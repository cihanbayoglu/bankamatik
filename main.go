package main

import (
	"bankamatik/models"
	"fmt"
)

func main() {
	var i int
	var j int
	k := models.Kisi{1, "Cihan", "Bayoğlu", true, nil}

	for i != -1 {
		fmt.Printf("\n*********************************\n")
		println("Yapmak istediğiniz işlemi giriniz")
		println("1-)  Hesap oluştur")
		println("2-)  Hesap bloke işlemleri")
		println("3-)  Para gönder")
		println("4-)  Para yatır")
		println("5-)  Para çek")
		println("6-)  Döviz al")
		println("7-)  Döviz bozdur")
		println("8-)  Hesap sil")
		println("9-)  Bakiye göster")
		println("10-) Hesap birleştirme")
		println("11-)Çıkış  için -1' giriniz.")
		fmt.Scan(&i)

		switch i {
		case 1:
			j++
			k.Hesaplar = append(k.Hesaplar, k.HesapOlustur(j))
			break
		case 2:
			if len(k.Hesaplar) == 0 {
				println("Bloke işlemleri yapmak için bir hesabınız bulunmamaktadır.")
				break
			}
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
			if len(k.Hesaplar) < 2 {
				println("Para gönderme işlemi için en az 2 hesaba ihtiyacınız vardır.")
				break
			}
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
		case 4:
			if len(k.Hesaplar) == 0 {
				println("Para yatırılacak bir hesabınız bulunmamaktadır.")
				break
			}

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
		case 5: //para çek
			if len(k.Hesaplar) == 0 {
				println("Para çekilecek bir hesabınız bulunmamaktadır.")
				break
			}
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
		case 6: //Döviz al (tl hesabından usd veya eur hesabına {para gönder : case3} )
			if len(k.Hesaplar) < 2 {
				println("Döviz alma işlemi için en az 1 adet tl 1 adet döviz hesabına ihtiyacınız vardır.")
				break
			}
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
		case 7: //Döviz bozdur (usd veya eur hesabından tl hesabona {para gönder: case3} )
			if len(k.Hesaplar) < 2 {
				println("Döviz alma işlemi için en az 1 adet tl 1 adet döviz hesabına ihtiyacınız vardır.")
				break
			}
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
		case 8:
			if len(k.Hesaplar) == 0 {
				println("Silinecek bir hesabınız bulunmamaktadır.")
				break
			}
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
		case 9:
			var e int // var olan hesapların id'lerini görmek için kullanılacak.
			for e = 0; e < len(k.Hesaplar); e++ {
				fmt.Printf("%d. Hesap , ID: %d , Durum: %t , Tür: %d , Bakiye: %.2f\n",
					e, k.Hesaplar[e].ID, k.Hesaplar[e].Durum, k.Hesaplar[e].Tur, k.Hesaplar[e].Bakiye)
			}
			break
		case 10:
			if len(k.Hesaplar) < 2 {
				println("Hesap birleştirme işlemi için en az 2 hesaba ihtiyacınız vardır.")
				break
			}
			var e, e2 int // var olan hesapların id'lerini görmek için kullanılacak.
			for e = 0; e < len(k.Hesaplar); e++ {
				fmt.Printf("%d. Hesap , ID: %d , Durum: %t , Tür: %d , Bakiye: %.2f\n",
					e, k.Hesaplar[e].ID, k.Hesaplar[e].Durum, k.Hesaplar[e].Tur, k.Hesaplar[e].Bakiye)
			}
			println("ana hesabı giriniz")
			fmt.Scan(&e)
			println("İçeriği ana hesaba aktarılacak hesabı giriniz. (Bu hesap işlem sonunda silinecektir.)")
			fmt.Scan(&e2)
			k.HesapBirlestir(&k.Hesaplar[e], &k.Hesaplar[e2])
			k.Hesaplar = append(k.Hesaplar[:e2], k.Hesaplar[e2+1:]...)
			break
		default:
			break
		}
	}

}
