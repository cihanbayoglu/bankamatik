package models

import "fmt"

const (
	usd = 9.3
	eur = 10.3
)

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

func (k Kisi) HesapOlustur(hesapID int, hesapTuru int) {
	hesap := Hesap{ID: hesapID, Sahibi: k, Tur: hesapTuru, Bakiye: 0, Durum: true}
	k.Hesaplar = append(k.Hesaplar, hesap)
	fmt.Println("Hesap Oluşturuldu.")

}

func (k Kisi) HesapBlokaj(hesap *Hesap, opsiyon bool) {
	if !opsiyon { // bloke etme işlemi için çalışır.
		if !hesap.Durum {
			fmt.Println("Seçmiş olduğunuz hesap halihazırda bloke.")
		} else {
			hesap.Durum = false
			fmt.Printf("ID: %d , Durum: %t , Tür: %d , Bakiye: %.2f\nHesap bloke edilmiştir."+
				" Artık bu hesapla işlem yapamazsınız.\n", hesap.ID, hesap.Durum, hesap.Tur, hesap.Bakiye)
		}
	} else { // bloke kaldırma işlemi için çalışır.
		if hesap.Durum {
			fmt.Println("Seçmiş olduğunuz hesap halihazırda bloke değil.")
		} else {
			hesap.Durum = true
			fmt.Printf("ID: %d , Durum: %t , Tür: %d , Bakiye: %.2f\nBloke kaldırılmıştır."+
				" Artık bu hesapla işlem yapabilirsiniz.\n", hesap.ID, hesap.Durum, hesap.Tur, hesap.Bakiye)
		}
	}

}

func (k Kisi) ParaGonder(alici *Hesap, gonderici *Hesap, tutar float64) {
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
	}
}

func (k Kisi) ParaYatir(hesap *Hesap, tutar float64) {
	if !hesap.Durum {
		fmt.Println("Bloke hesap ile herhangi bir işlem yapılamaz.")

	} else {
		hesap.Bakiye += tutar
		fmt.Println("İşleminiz başarıyla gerçekleşmiştir.")
	}
}

func (k Kisi) ParaCek(hesap *Hesap, tutar float64) {
	if !hesap.Durum {
		fmt.Println("Bloke hesap ile herhangi bir işlem yapılamaz.")

	} else {
		hesap.Bakiye -= tutar
		fmt.Println("İşleminiz başarıyla gerçekleşmiştir.")
	}
}

func (k Kisi) DovizAl(doviz *Hesap, tl *Hesap, tutar float64) {
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

	} else if doviz.Tur == 3 {
		tl.Bakiye -= tutar
		doviz.Bakiye += tutar / eur
		fmt.Printf("%.1f kurundan euro alma işlemi başarıyla tamamlanmıştır.\n"+
			" Tl hesabı bakiye : %.2f, euro hesabı bakiye: %.2f\n", eur, tl.Bakiye, doviz.Bakiye)
	}
}

func (k Kisi) DovizBoz(doviz *Hesap, tl *Hesap, tutar float64) {
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
	} else if doviz.Tur == 3 {
		doviz.Bakiye -= tutar
		tl.Bakiye += tutar * eur
		fmt.Printf(" %.1f kurundan euro bozdurma işlemi başarıyla tamamlanmıştır.\n"+
			" Tl hesabı bakiye : %.2f, euro hesabı bakiye: %.2f\n", usd, tl.Bakiye, doviz.Bakiye)
	}
}

func (k Kisi) HesapSil(HesapID int) Hesap {
	hesap := Hesap{ID: HesapID, Sahibi: k, Bakiye: 0, Durum: true}
	fmt.Println("Oluşturmak istediğiniz hesap türünü seciniz")
	fmt.Println("1-)tl -- 2-)dolar -- 3-)euro")
	fmt.Scan(&hesap.Tur)

	return hesap
}

func (k Kisi) HesapBirlestir(ana *Hesap, silinecek *Hesap, e2 int) {
	if ana.Tur != silinecek.Tur {
		fmt.Println("Sadece aynı tür hesapları birleştirebilirsiniz.")

	} else if !ana.Durum || !silinecek.Durum {
		fmt.Println("Bloke hesaplar ile işlem yapılamaz.")
	} else {
		ana.Bakiye += silinecek.Bakiye
		fmt.Println("Hesap birleştirme işlemi başarıyla gerçekleştirildi. Ana hesap :")
		fmt.Printf(" ID: %d , Durum: %t , Tür: %d , Bakiye: %.2f\n", ana.ID, ana.Durum, ana.Tur, ana.Bakiye)
		k.Hesaplar = append(k.Hesaplar[:e2], k.Hesaplar[e2+1:]...)
	}

}
