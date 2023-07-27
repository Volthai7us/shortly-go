# URL Shortener

URL Shortener, gelen uzun URL'leri kısa URL'lere dönüştürmek için bir Go uygulamasıdır. Bu uygulama DDD (Domain Driven Design) prensiplerine göre oluşturulmuştur ve Kubernetes üzerinden servis edilebilir.

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
  /deployment.yaml
  /service.yaml
  /run.sh
  /build.sh
  /README.md
  /go.mod
```

---

## Nasıl Çalışır

Bu uygulama, bir HTTP POST isteği yoluyla alınan herhangi bir URL'yi kısaltır.

1. Kullanıcı, kısaltılacak bir URL'yi HTTP POST isteği ile `/create` endpoint'ine gönderir.
2. Bu istek `url_controller` tarafından işlenir ve `url_shortener_service`'e aktarılır.
3. `url_shortener_service`, URL'yi önce `url_repository`'de arar. Eğer URL daha önce kısaltılmışsa, var olan kısa URL döndürülür. Eğer URL daha önce kısaltılmamışsa, yeni bir kısa URL oluşturulur ve bu URL `url_repository`'e kaydedilir.
4. Oluşturulan kısa URL, HTTP yanıtında kullanıcıya geri gönderilir.
5. Kullanıcı, `/` endpoint'ine kısa URL'yi göndererek orijinal URL'ye yönlendirilir.

- Kubernetes replikaları arasında tutarlılık sağlamak için, sessionAffinity ClientIP olarak ayarlanmıştır.
- Bu sayede, IP adresinden gelen isteklerde aynı pod'a yönlendirilir.

---

## Kurulum

Bu projeyi çalıştırmak için aşağıdaki adımları takip edebilirsiniz:

1. Bu repo'yu klonlayın:
   ```bash
   git clone https://github.com/yourusername/url-shortener.git
   cd url-shortener
   ```
2. Uygulamayı build edin ve çalıştırın:
   ```bash
   go build ./cmd/url-shortener
   ./url-shortener
   ```
3. Artık uygulama `http://localhost:5173` adresinde çalışıyor olmalıdır.

---

## Testler

Projenin testlerini çalıştırmak için aşağıdaki komutu kullanabilirsiniz:

```bash
go test test/url_shortener_service_test.go
```

---

## Docker

Bu projeyi Docker üzerinde çalıştırmak için, aşağıdaki komutları kullanabilirsiniz:

1. Docker image'ini build edin:
   ```bash
   docker build -t url-shortener .
   ```
2. Docker container'ı başlatın:
   ```bash
   docker run -p 5173:5173 url-shortener
   ```

---

## Kubernetes

Bu projeyi Kubernetes üzerinde çalıştırmak için, aşağıdaki komutları kullanabilirsiniz:

1. Docker image'ini build edin ve Docker Hub'a push edin:
   ```bash
   docker build -t yourusername/url-shortener .
   docker push yourusername/url-shortener
   ```
2. Kubernetes deployment ve service'ini apply edin:
   ```bash
   kubectl apply -f deployment.yaml
   kubectl apply -f service.yaml
   ```
