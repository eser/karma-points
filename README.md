# karma-points DevDoc

## Begginer's Guide 🏁

### Introduction

- **Projenin Amacı ve Genel Tanıtımı:**
  Karma Points, kişilerin birbirlerine yaptıkları takdirleri iletebilmeleri ve bu takdirler üzerinden rozetler kazanabilecekleri bir ödüllendirme sistemidir. Uygulama, kullanıcıların olumlu davranışlarını teşvik eder ve topluluk içinde pozitif bir atmosfer yaratır. Takdirler, belirli etkinlikler ve başarılar üzerine verilebilir ve bu sayede kullanıcılar, toplulukları içinde tanınma ve motivasyon kazanırlar.

- **Projenin Hedef Kitlesi ve Çözdüğü Problemler:**
  Karma-Points, özellikle kurumsal şirketler, eğitim kurumları ve büyük topluluklar için tasarlanmıştır. Hedef kitlesi, takım çalışmasını teşvik etmek, çalışan veya öğrenci motivasyonunu artırmak isteyen organizasyonlardır. Uygulama, kişisel çabaların fark edilmemesi, düşük motivasyon ve takım içi iletişim sorunları gibi problemleri çözmeyi amaçlar. Karma-Points, kullanıcıların birbirlerine verdikleri takdirler ile daha güçlü bir bağ kurmalarına ve daha verimli çalışmalara olanak tanır.

### Installation Instructions

- **Gerekli Ön Koşullar:**
  - Node.js (versiyon X.X.X ve üzeri)
  - Git

- **Clone Aşamaları:**

  ```bash
  git clone https://github.com/eser/karma-points.git
  cd karma-points
  ```

- **Clone aşamaları:**

  ```bash
  npm install
  npm run dev
  ```

### Getting Started

örnek yapılandırma .env dosyaları

```bash
DB_HOST=localhost
DB_USER=root
DB_PASS=password
  ```

## Features and Usage 🗝️

### Key Features

- Kullanıcıların birbirlerine takdir gönderebilmesi.
- Rozet yönetimi.
- Etkinlik tabanlı ödüllendirme sistemi.

### step-by-step usage

nasıl kullanılacağına dair detaylı açıklamalar ve usa cases

### Commands and Params

eğer cli ise komut ve parametre açıklamaları

## Contribution Guide 🤝

### Contribution Process

Katkıda bulunmak isteyenler için adım adım talimatlar:

1. **Proje Fork'lama:**
   - GitHub hesabınıza giriş yapın.
   - Karma-Points projesinin GitHub sayfasına gidin.
   - Sağ üst köşedeki "Fork" butonuna tıklayarak projeyi kendi hesabınıza fork'layın.

2. **Yeni Bir Branch Oluşturma:**

   - Yeni bir branch oluşturun ve bu branch'e geçin:

     ```bash
     git checkout -b yeni-ozellik
     ```

3. **Değişiklikleri Yapma:**
   - Kendi branch'inizde gerekli değişiklikleri yapın.
   - Değişikliklerinizi commit edin:

     ```bash
     git add .
     git commit -m 'Yeni özellik ekle'
     ```

4. **Branch'i Push'lama:**
   - Değişikliklerinizi uzak depoya gönderin:

     ```bash
     git push origin yeni-ozellik
     ```

5. **Pull Request Açma:**
   - GitHub sayfanıza gidin ve fork'ladığınız repository'yi açın.
   - "New Pull Request" butonuna tıklayarak yeni bir pull request oluşturun.
   - Yapılan değişikliklerin detaylarını açıklayan bir açıklama yazın ve pull request'i gönderin.

### Coding Standards

Proje için geçerli olan kodlama standartları ve en iyi uygulamalar:

- **Kodlama Stili:**
  - Kod yazarken anlaşılır ve tutarlı bir stil kullanın.
  - Kodunuzun okunabilir olmasına dikkat edin ve açıklayıcı yorumlar ekleyin.
  - Değişken ve fonksiyon isimlendirmelerinde anlamlı ve açıklayıcı isimler kullanın.

- **Linting ve Formatlama:**
  - Projede kullanılan linter araçlarını kullanarak kodunuzu kontrol edin.
  - Kod formatlama kurallarına uyun ve gerekli formatlama araçlarını kullanarak kodunuzu formatlayın.

- **Fonksiyon ve Modül Tasarımı:**
  - Fonksiyonlarınızı ve modüllerinizi küçük, anlaşılır ve tek bir iş yapacak şekilde tasarlayın.
  - Kod tekrarından kaçının ve yeniden kullanılabilir kod yazmaya özen gösterin.

### Tests and Verifications

Testlerin nasıl yazılacağı ve doğrulama süreçleri:

- **Test Yazma:**
  - Proje için belirlenen test çerçevesini kullanarak testlerinizi yazın (örneğin, Jest, Mocha).
  - Birim testleri ve entegrasyon testleri yazarak kodunuzun doğruluğunu kontrol edin.
  - Test kapsamını artırmak için her yeni özellik veya değişiklik için test yazmayı ihmal etmeyin.

- **Test Çalıştırma:**
  - Testleri çalıştırmak için proje kök dizininde aşağıdaki komutu kullanın:

    ```bash
    npm test
    ```

- **Doğrulama ve İnceleme:**
  - Değişikliklerinizi göndermeden önce tüm testlerin geçtiğinden emin olun.
  - Kodunuzun gözden geçirilmesi için pull request açın ve diğer geliştiricilerin incelemesine sunun.
  - Gözden geçirme sırasında yapılan geri bildirimleri dikkate alarak gerekli düzeltmeleri yapın.

## FAQ 🔍

### Common Questions and Issues

- Sıkça sorulan sorular ve genel sorunlar.

### Troubleshooting

- Hata ayıklama ve sorun giderme ipuçları.
