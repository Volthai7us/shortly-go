# URL Shortener

URL Shortener, gelen uzun URL'leri kısa URL'lere dönüştürmek için bir Go uygulamasıdır. Bu uygulama DDD (Domain Driven Design) prensiplerine göre oluşturulmuştur ve Kubernetes üzerinden servis edilebilir.

(https://www.youtube.com/watch?v=650VEfIOD1s)

---

## Yapı

Bu projenin yapılanması aşağıdaki gibi organize edilmiştir:

```plaintext
/url-shortener
  /cmd
    /url-shortener
      main.go
  /app
    /controllers
      url_controller.go
    /entities
      url.go
    /helper
      helper.go
    /repositories
      url_repository.go
    /services
      url_shortener_service.go
  /test
    url_shortener_test.go
  /Dockerfile
```

---

## Nasıl Çalışır

Bu uygulama, bir HTTP POST isteği yoluyla alınan herhangi bir URL'yi kısaltır.

1. Kullanıcı, kısaltılacak bir URL'yi HTTP POST isteği ile `/create` endpoint'ine gönderir.
2. Bu istek `url_controller` tarafından işlenir ve `url_shortener_service`'e aktarılır.
3. `url_shortener_service`, URL'yi önce `url_repository`'de arar. Eğer URL daha önce kısaltılmışsa, var olan kısa URL döndürülür. Eğer URL daha önce kısaltılmamışsa, yeni bir kısa URL oluşturulur ve bu URL `url_repository`'e kaydedilir.
4. Oluşturulan kısa URL, HTTP yanıtında kullanıcıya geri gönderilir.
5. Kullanıcı, `/` endpoint'ine kısa URL'yi göndererek orijinal URL'ye yönlendirilir.
6. SHA256 algoritması kullanılarak oluşturulan kısa URL, 10 karakter uzunluğundadır.

- Kubernetes replikaları arasında tutarlılık sağlamak için, PersistentVolume ve PersistentVolumeClaim kullanılmıştır.
- Bu sayede, uygulama herhangi bir Kubernetes pod'undan kaldırılsa bile, veriler kaybolmaz.
- Ve uygulama tekrar başlatıldığında, veriler PersistentVolume'dan geri yüklenir.

---

## Kurulum

Bu projeyi çalıştırmak için aşağıdaki adımları takip edebilirsiniz:

1. Bu repo'yu klonlayın:
   ```bash
   git clone https://github.com/yourusername/url-shortener.git
   cd url-shortener
   ```
2. Docker image'ini build edin ve Docker Hub'a push edin ve Kubernetes yaml dosyalarını apply edin:
   ```bash
   ./build.sh
   ```
3. Uygulamayı Kubernetes üzerinde çalıştırın:
   ```bash
    ./run.sh
   ```

---

## Testler

Projenin testlerini çalıştırmak için aşağıdaki komutu kullanabilirsiniz:

```bash
go test test/url_shortener_service_test.go
```
