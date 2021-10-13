# bankamatik
#### Yapılacaklar

CLI tabanlı kullanıcı banka hesabı açabildiği, para yatırabildiği, para çekebildiği ve Döviz işlemleri yapabildiği (sadece dolar ve euro) olacak. (Not Dolar kurları sabit hard coded girilecek.) Herhangi bir finansal hatayı engelleyecek ve hile yapılmasına engel olunacak tüm testler eklenmiş olmalıdır.

Kullanıcı Her işlem sonrasında menüye atılmalı ve işlem seçimide 0 seçerse çıkış yapabilmelidir. Çıkış işlemi gerçekleşirken tüm veriler a.txt adlı bir dosyada tutulmaldır.

- Hesap

``
type Hesap struct {
    ID int    
    Sahibi *Kisi
    HesapTuru int
    Bakiye float64
    Durum bool // false ise bloke hesap olarak hiç bir işlem gerçekleştirememesi gerekmektedir.
}``

- Kisi

``
type Kisi struct {
ID int    
Ad string
Soyad string
Hesaplar []Hesap
Durum bool
}``

#### Fonksiyonlar

- Hesap oluştur
- Hesap bloke et (Kullanıcı veya  Hesap bloke edilebilsin iki blokeden biri varsa işlem yapılmasına engel olsun)
- Para gönder
- Para Yatır
- Para çek
- Döviz al
- Döviz bozdur
- Hesap sil (Bloke ise silinemez)
- Bakiye göster
- Hesap birleştirme (Kullanıcı aynı parabiriminden farklı hesaplar açtıysa bu hesapları birleştirebilme imkanı oluşsun. Dövizlerde Kur farklı hesaplanıp farklı para birimleri ile birleştirilebilsin.)

- (Opsiyonel olarak) Kullanıcı İşlem kayıtları islemkayit.log dosyası altında kayıt altında tutulmalıdır.
#### Kullanılacaklar

- github'dan proje oluşturulup go mod yapılmış olmalı
