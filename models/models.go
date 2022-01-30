package models

import "fmt"

const (
	usd = 9.3
	eur = 10.3
)

var hesapId int //Hs

//type Hesap struct { ID int Sahibi *Kisi HesapTuru int Bakiye float64 Durum bool
//false ise bloke hesap olarak hiç bir işlem gerçekleştirememesi gerekmektedir. }
type Hesap struct {
	ID     int
	Sahibi Kisi
	Tur    int
	Bakiye float64
	Durum  bool
}

//type Kisi struct { ID int Ad string Soyad string Hesaplar []Hesap Durum bool }

type Kisi struct {
	ID       int
	Ad       string
	Soyad    string
	Durum    bool
	Hesaplar []Hesap
}

func HataliMi(kisi *Kisi, input ...int) bool {
	for i := range input {
		if input[i] > len(kisi.Hesaplar)-1 || input[i] < 0 {

			return true
		}
	}
	return false
}

func (k Kisi) HesapOlustur(kisi *Kisi) bool {
	var hesapTuru int
	fmt.Println("Oluşturmak istediğiniz hesap türünü seciniz")
	fmt.Println("1-)tl -- 2-)dolar -- 3-)euro")
	fmt.Scan(&hesapTuru)
	if hesapTuru >= 1 && hesapTuru <= 3 {
		hesapId++
		hesap := Hesap{ID: hesapId, Sahibi: *kisi, Tur: hesapTuru, Bakiye: 0, Durum: true}
		kisi.Hesaplar = append(kisi.Hesaplar, hesap)
		fmt.Println("Hesap Oluşturuldu.")
		return true
	} else {
		fmt.Println("Hatalı hesap türü girdiğiniz için hesap oluşturulamadı.")
		return false
	}

}

func (k Kisi) HesapBlokaj() bool {
	var hesapNo int
	var opsiyon bool
	println("Blokaj işlemi yapılacak hesabı giriniz.")
	fmt.Scan(&hesapNo)
	println("Bloke etmek için 0, blokeyi kaldırmak için 1 giriniz.")
	fmt.Scan(&opsiyon)

	if HataliMi(&k, hesapNo) {
		println("Yanlış hesap girdiniz.")
		return false
	} else {
		if !opsiyon { // bloke etme işlemi için çalışır.
			if !k.Hesaplar[hesapNo].Durum {
				fmt.Println("Seçmiş olduğunuz hesap halihazırda bloke.")
			} else {
				k.Hesaplar[hesapNo].Durum = false
				fmt.Printf("ID: %d , Durum: %t , Tür: %d , Bakiye: %.2f\nHesap bloke edilmiştir."+
					" Artık bu hesapla işlem yapamazsınız.\n", k.Hesaplar[hesapNo].ID, k.Hesaplar[hesapNo].Durum, k.Hesaplar[hesapNo].Tur, k.Hesaplar[hesapNo].Bakiye)
			}
		} else { // bloke kaldırma işlemi için çalışır.
			if k.Hesaplar[hesapNo].Durum {
				fmt.Println("Seçmiş olduğunuz hesap halihazırda bloke değil.")
			} else {
				k.Hesaplar[hesapNo].Durum = true
				fmt.Printf("ID: %d , Durum: %t , Tür: %d , Bakiye: %.2f\nBloke kaldırılmıştır."+
					" Artık bu hesapla işlem yapabilirsiniz.\n", k.Hesaplar[hesapNo].ID, k.Hesaplar[hesapNo].Durum, k.Hesaplar[hesapNo].Tur, k.Hesaplar[hesapNo].Bakiye)
			}
		}
	}
	return true
}

func (k Kisi) ParaGonder(alici *Hesap, gonderici *Hesap, tutar float64) bool {
	if alici.Durum == false && gonderici.Durum == false {
		fmt.Println("Bloke hesap ile herhangi bir işlem yapılamaz.")

	} else if tutar > gonderici.Bakiye {
		fmt.Println("Hesabınızda yeterli bakiye bulunmamaktadır.")

	} else if alici.Tur != gonderici.Tur {
		fmt.Println("Farklı tür hesaplar arası para gönderilemez.(anamenüde 6 veya 7 opsiyonunu seçiniz.")

	} else {
		gonderici.Bakiye -= tutar
		alici.Bakiye += tutar
		fmt.Println("İşleminiz başarıyla gerçekleşmiştir.")
		return true
	}
	return false
}

func (k Kisi) ParaYatir(hesap *Hesap, tutar float64) bool {
	if !hesap.Durum {
		fmt.Println("Bloke hesap ile herhangi bir işlem yapılamaz.")

	} else {
		hesap.Bakiye += tutar
		fmt.Println("İşleminiz başarıyla gerçekleşmiştir.")
		return true
	}
	return false
}

func (k Kisi) ParaCek(hesap *Hesap, tutar float64) bool {
	if !hesap.Durum {
		fmt.Println("Bloke hesap ile herhangi bir işlem yapılamaz.")

	} else if tutar > hesap.Bakiye {
		fmt.Println("Hesabınızda yeterli bakiye bulunmamaktadır.")
	} else {
		hesap.Bakiye -= tutar
		fmt.Println("İşleminiz başarıyla gerçekleşmiştir.")
		return true
	}
	return false
}

func (k Kisi) DovizAl(doviz *Hesap, tl *Hesap, tutar float64) bool {
	if !doviz.Durum || !tl.Durum {
		fmt.Println("Bloke hesap ile herhangi bir işlem yapılamaz.")

	} else if doviz.Tur == 1 {
		fmt.Println("Tl hesabına döviz ekleyemezsiniz.")

	} else if tl.Tur != 1 {
		fmt.Println("Sadece tl hesabından döviz alabilirsiniz.")

	} else if tutar > tl.Bakiye {
		fmt.Println("Hesabınızda yeterli bakiye bulunmamaktadır.")

	} else if doviz.Tur == 2 {
		tl.Bakiye -= tutar
		doviz.Bakiye += tutar / usd
		fmt.Printf(" %.1f kurundan dolar alma işlemi başarıyla tamamlanmıştır.\n"+
			" Tl hesabı bakiye : %.2f, döviz hesabı bakiye: %.2f\n", usd, tl.Bakiye, doviz.Bakiye)
		return true

	} else if doviz.Tur == 3 {
		tl.Bakiye -= tutar
		doviz.Bakiye += tutar / eur
		fmt.Printf("%.1f kurundan euro alma işlemi başarıyla tamamlanmıştır.\n"+
			" Tl hesabı bakiye : %.2f, euro hesabı bakiye: %.2f\n", eur, tl.Bakiye, doviz.Bakiye)
		return true
	}
	return false
}

func (k Kisi) DovizBoz(doviz *Hesap, tl *Hesap, tutar float64) bool {
	if !doviz.Durum || !tl.Durum {
		fmt.Println("Bloke hesap ile herhangi bir işlem yapılamaz.")

	} else if doviz.Tur == 1 {
		fmt.Println("Tl hesabından döviz bozduramazsınız.")

	} else if tl.Tur != 1 {
		fmt.Println("Bozdurduğunuz dövizi sadece tl hesabına aktarabilirsiniz.")

	} else if tutar > doviz.Bakiye {
		fmt.Println("Hesabınızda yeterli bakiye bulunmamaktadır.")

	} else if doviz.Tur == 2 {
		doviz.Bakiye -= tutar
		tl.Bakiye += tutar * usd
		fmt.Printf(" %.1f kurundan dolar bozdurma işlemi başarıyla tamamlanmıştır.\n"+
			" Tl hesabı bakiye : %.2f, dolar hesabı bakiye: %.2f\n", usd, tl.Bakiye, doviz.Bakiye)
		return true
	} else if doviz.Tur == 3 {
		doviz.Bakiye -= tutar
		tl.Bakiye += tutar * eur
		fmt.Printf(" %.1f kurundan euro bozdurma işlemi başarıyla tamamlanmıştır.\n"+
			" Tl hesabı bakiye : %.2f, euro hesabı bakiye: %.2f\n", usd, tl.Bakiye, doviz.Bakiye)
		return true
	}
	return false
}

/*func (k Kisi) HesapSil(HesapID int)  {
	println("Silmek istediğiniz hesabı giriniz")
	fmt.Scan(&e)
	if !k.Hesaplar[e].Durum {
		println("Bloke hesap silinemez.")

	} else {
		k.Hesaplar = append(k.Hesaplar[:e], k.Hesaplar[e+1:]...)
		fmt.Printf("%d. hesabınız silindi.", e)
	}
}*/

func (k Kisi) HesapBirlestir(ana *Hesap, silinecek *Hesap, e2 int) bool {
	if ana.Tur != silinecek.Tur {
		fmt.Println("Sadece aynı tür hesapları birleştirebilirsiniz.")

	} else if !ana.Durum || !silinecek.Durum {
		fmt.Println("Bloke hesaplar ile işlem yapılamaz.")
	} else {
		ana.Bakiye += silinecek.Bakiye
		fmt.Println("Hesap birleştirme işlemi başarıyla gerçekleştirildi. Ana hesap :")
		fmt.Printf(" ID: %d , Durum: %t , Tür: %d , Bakiye: %.2f\n", ana.ID, ana.Durum, ana.Tur, ana.Bakiye)
		k.Hesaplar = append(k.Hesaplar[:e2], k.Hesaplar[e2+1:]...)
		return true
	}

	return false
}

func (k Kisi) BakiyeGoster() bool {
	for i := 0; i < len(k.Hesaplar); i++ {
		fmt.Printf("%d. Hesap , ID: %d , Durum: %t , Tür: %d , Bakiye: %.2f\n",
			i, k.Hesaplar[i].ID, k.Hesaplar[i].Durum, k.Hesaplar[i].Tur, k.Hesaplar[i].Bakiye)
	}
	return true
}
